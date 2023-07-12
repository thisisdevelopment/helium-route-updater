package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_api "github.com/thisisdevelopment/helium-route-updater/pkg/helium-api"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/lns"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
	"io"
	"log"
	"strconv"
	"time"
)

func main() {
	config := types.ConfigFromEnv()
	ch := make(chan types.DeviceEvent)
	keypair := helium_crypto.KeyPairFromString(config.Helium.KeyPair)
	heliumClient := helium_api.NewClient(config.Helium.Server, keypair)
	lnsClient := lns.NewClient(config.Lns)

	go helium_api.Run(heliumClient, config.Helium.RouteId, ch)
	go scheduledSync(config.Helium.RouteId, ch, heliumClient, lnsClient)

	lnsClient.Listen(ch)
}

func scheduledSync(routeId string, ch chan<- types.DeviceEvent, heliumClient *helium_api.Client, lnsClient lns.Client) {
	sync(routeId, ch, heliumClient, lnsClient)
	for _ = range time.Tick(1 * time.Hour) {
		sync(routeId, ch, heliumClient, lnsClient)
	}
}

func sync(routeId string, ch chan<- types.DeviceEvent, heliumClient *helium_api.Client, lnsClient lns.Client) {
	devices := lnsClient.GetDevices()
	ctx := context.Background()
	euiStream, err := heliumClient.NewRouteClient().GetEuis(ctx, helium_crypto.SignRequest(&iot_config.RouteGetEuisReqV1{RouteId: routeId}, heliumClient.Keypair))
	if err != nil {
		log.Fatalf("cannot receive %v", err)
	}

	euis := []*types.Device{}
	for {
		resp, err := euiStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		euis = append(euis, &types.Device{
			DevEui:  resp.DevEui,
			JoinEui: resp.AppEui,
		})
	}

	devicesEuis := append(devices)
	ch <- syncDevices(devicesEuis, euis, "euis")

	skfs := []*types.Device{}
	skfStream, err := heliumClient.NewRouteClient().GetSkfs(ctx, helium_crypto.SignRequest(&iot_config.RouteSkfGetReqV1{RouteId: routeId}, heliumClient.Keypair))
	for {
		resp, err := skfStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("cannot receive %v", err)
		}
		sk, _ := hex.DecodeString(resp.SessionKey)
		skfs = append(skfs, &types.Device{
			DevAddr:    resp.Devaddr,
			SessionKey: sk,
		})
	}
	devicesSkfs := append(devices)
	ch <- syncDevices(devicesSkfs, skfs, "skfs")
}

func syncDevices(lns []*types.Device, helium []*types.Device, syncType string) types.DeviceEvent {
	update := []*types.Device{}
	del := []*types.Device{}

	existing := map[string]*types.Device{}
	for _, dev := range lns {
		key := ""
		if syncType == "euis" {
			key = strconv.FormatUint(dev.DevEui, 10)
		} else {
			key = fmt.Sprintf("%d:%s", dev.DevAddr, hex.EncodeToString(dev.SessionKey))
		}
		existing[key] = dev
	}

	for _, dev := range helium {
		key := ""
		if syncType == "euis" {
			key = strconv.FormatUint(dev.DevEui, 10)
		} else {
			key = fmt.Sprintf("%d:%s", dev.DevAddr, hex.EncodeToString(dev.SessionKey))
		}

		match, ok := existing[key]
		if !ok {
			del = append(del, dev)
		} else {
			if syncType == "euis" && match.JoinEui != dev.JoinEui {
				//for syncType == "skfs" the key contains all data, so no need to update if we find a match
				//for syncType == "euis" we only need to update if the joinEui is updated
				update = append(update, dev)
			}
			delete(existing, key)
		}
	}

	for _, dev := range existing {
		var d *types.Device
		if syncType == "euis" {
			d = &types.Device{JoinEui: dev.JoinEui, DevEui: dev.DevEui}
		} else {
			d = &types.Device{DevAddr: dev.DevAddr, SessionKey: dev.SessionKey}
		}
		update = append(update, d)
	}

	return types.DeviceEvent{
		Update: update,
		Delete: del,
	}
}
