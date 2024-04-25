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
	"os"
	"strings"
	"time"
)

func main() {
	var err error
	var keypair *helium_crypto.KeyPair

	config := types.ConfigFromEnv()
	ch := make(chan types.DeviceEvent)
	syncCh := make(chan bool)

	if strings.HasPrefix(config.Helium.KeyPair, "/") || strings.HasPrefix(config.Helium.KeyPair, "./") {
		dat, err := os.ReadFile(config.Helium.KeyPair)
		if err != nil {
			log.Fatalf("Unable to read keypair from file: %s\n", config.Helium.KeyPair)
		}

		keypair, err = helium_crypto.KeyPairFromBytes(dat)
		if err != nil {
			log.Fatalf("Unable to parse keypair from file: %s\n", config.Helium.KeyPair)
		}

	} else {
		keypair, err = helium_crypto.KeyPairFromString(config.Helium.KeyPair)
		if err != nil {
			log.Fatalf("Unable to parse keypair from string: %s\n", config.Helium.KeyPair)
		}
	}

	heliumClient := helium_api.NewClient(config.Helium.Server, keypair)
	lnsClient := lns.NewClient(config.Lns)

	fmt.Printf("Starting\n")

	go helium_api.Run(heliumClient, config.Helium.RouteId, ch)
	go lnsClient.Listen(ch)
	go sync(config.Helium.RouteId, ch, syncCh, heliumClient, lnsClient)
	scheduledSync(syncCh)
}

func scheduledSync(syncCh chan<- bool) {
	syncCh <- true
	for range time.Tick(1 * time.Hour) {
		syncCh <- true
	}
}

func sync(routeId string, ch chan<- types.DeviceEvent, syncCh <-chan bool, heliumClient *helium_api.Client, lnsClient lns.Client) {
	ctx := context.Background()
	for range syncCh {
		fmt.Printf("[sync] starting full sync\n")

		euiStream, err := heliumClient.NewRouteClient().GetEuis(ctx, helium_crypto.SignRequest(&iot_config.RouteGetEuisReqV1{RouteId: routeId}, heliumClient.Keypair))
		if err != nil {
			log.Fatalf("cannot get euis %v", err)
		}

		euis := []*types.Device{}
		for {
			resp, err := euiStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("cannot receive eui %v", err)
			}
			euis = append(euis, &types.Device{
				DevEui:  resp.DevEui,
				JoinEui: resp.AppEui,
			})
		}
		fmt.Printf("[sync] Found %d euis in our helium route\n", len(euis))

		skfs := []*types.Device{}
		skfStream, err := heliumClient.NewRouteClient().ListSkfs(ctx, helium_crypto.SignRequest(&iot_config.RouteSkfListReqV1{RouteId: routeId}, heliumClient.Keypair))
		if err != nil {
			log.Fatalf("cannot get skfs %v", err)
		}
		for {
			resp, err := skfStream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("cannot receive skf %v", err)
			}
			sk, _ := hex.DecodeString(resp.SessionKey)

			if len(sk) > 0 {
				skfs = append(skfs, &types.Device{
					DevAddr:    resp.Devaddr,
					SessionKey: sk,
				})
			} else {
				fmt.Printf("[sync] uhhm .. skf does not contain session key? %#V\n", resp)
			}
		}
		fmt.Printf("[sync] Found %d skfs in our helium route\n", len(skfs))

		devices := lnsClient.GetDevices()
		fmt.Printf("[sync] Found %d devices in the lns\n", len(devices))

		//TODO: we have race condition here .. while we are retrieving the devices from helium/lns we can receive updates
		//      these updates will (possibly) not be reflected here .. leading to false diffs and bogus updates :-/
		//      ideally we disable the updates during a full sync; but that is currently not possible as the full
		//      sync can take a long time with a lot of devices.
		syncDevices(ch, devices, euis, "euis")
		syncDevices(ch, devices, skfs, "skfs")
	}
}

func getDeviceKey(dev *types.Device, syncType string) string {
	key := ""
	if syncType == "euis" {
		key = fmt.Sprintf("%x", dev.DevEui)
	} else {
		key = fmt.Sprintf("%x:%s", dev.DevAddr, hex.EncodeToString(dev.SessionKey))
	}
	return key
}

func syncDevices(ch chan<- types.DeviceEvent, lns []*types.Device, helium []*types.Device, syncType string) {
	update := []*types.Device{}
	del := []*types.Device{}

	existing := map[string]*types.Device{}
	for _, dev := range lns {
		key := getDeviceKey(dev, syncType)
		existing[key] = dev
	}

	for _, dev := range helium {
		key := getDeviceKey(dev, syncType)
		match, ok := existing[key]
		if !ok {
			fmt.Printf("[sync] %s not found in lns => remove from helium\n", key)
			del = append(del, dev)
		} else {
			if syncType == "euis" && match.JoinEui != dev.JoinEui {
				//for syncType == "skfs" the key contains all data, so no need to update if we find a match
				//for syncType == "euis" we only need to update if the joinEui is updated
				fmt.Printf("[sync] found different joinEui for %s (%d vs %d) => update on helium\n", key, match.JoinEui, dev.JoinEui)
				update = append(update, dev)
			}
			delete(existing, key)
		}
	}

	for key, dev := range existing {
		var d *types.Device
		if syncType == "euis" {
			if dev.JoinEui != 0x0 {
				fmt.Printf("[sync] missing eui %s => adding on helium\n", key)
				d = &types.Device{JoinEui: dev.JoinEui, DevEui: dev.DevEui}
			} else {
				fmt.Printf("[sync] eui without joineui %s => not adding\n", key)
			}
		} else if len(dev.SessionKey) > 0 {
			fmt.Printf("[sync] missing sessionkey %s => adding on helium\n", key)
			d = &types.Device{DevAddr: dev.DevAddr, SessionKey: dev.SessionKey}
		}
		if d != nil {
			update = append(update, d)
		}
	}

	if len(update)+len(del) > 0 {
		ch <- types.DeviceEvent{
			Update: update,
			Delete: del,
		}
	} else {
		fmt.Printf("[sync] %s are in sync\n", syncType)
	}
}
