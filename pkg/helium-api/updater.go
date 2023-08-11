package helium_api

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

func Run(client *Client, routeId string, ch <-chan types.DeviceEvent) {
	ctx := context.Background()
	routeClient := client.NewRouteClient()

	for event := range ch {
		fmt.Printf("[helium] start transaction\n")

		euis := []*iot_config.RouteUpdateEuisReqV1{}
		skfs := []*iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{}
		for _, dev := range event.Update {
			fmt.Printf("[helium] update device: %#v\n", dev)
			if dev.DevAddr != 0 && len(dev.SessionKey) > 0 {
				fmt.Printf("[helium][%s] adding skf %x %s\n", routeId, dev.DevAddr, hex.EncodeToString(dev.SessionKey))
				skfs = append(skfs, &iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{
					Devaddr:    dev.DevAddr,
					SessionKey: hex.EncodeToString(dev.SessionKey),
					Action:     iot_config.ActionV1_add,
					MaxCopies:  dev.MaxCopies,
				})
			}
			if dev.JoinEui != 0 && dev.DevEui != 0 {
				fmt.Printf("[helium][%s] adding euipair %x %x\n", routeId, dev.DevEui, dev.JoinEui)
				euis = append(euis, &iot_config.RouteUpdateEuisReqV1{
					Action: iot_config.ActionV1_add,
					EuiPair: &iot_config.EuiPairV1{
						RouteId: routeId,
						AppEui:  dev.JoinEui,
						DevEui:  dev.DevEui,
					},
				})
			}
		}

		for _, dev := range event.Delete {
			if dev.DevAddr != 0 && len(dev.SessionKey) > 0 {
				fmt.Printf("[helium][%s] removing skf %x %s\n", routeId, dev.DevAddr, hex.EncodeToString(dev.SessionKey))
				skfs = append(skfs, &iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{
					Devaddr:    dev.DevAddr,
					SessionKey: hex.EncodeToString(dev.SessionKey),
					Action:     iot_config.ActionV1_remove,
				})
			}
			if dev.JoinEui != 0 && dev.DevEui != 0 {
				fmt.Printf("[helium][%s] removing euipair %x %x\n", routeId, dev.DevEui, dev.JoinEui)
				euis = append(euis, &iot_config.RouteUpdateEuisReqV1{
					Action: iot_config.ActionV1_remove,
					EuiPair: &iot_config.EuiPairV1{
						RouteId: routeId,
						AppEui:  dev.JoinEui,
						DevEui:  dev.DevEui,
					},
				})
			}
		}

		if len(euis) > 0 {
			u, err := routeClient.UpdateEuis(ctx)
			if err != nil {
				panic(err)
			}
			for _, eui := range euis {
				err = u.Send(helium_crypto.SignRequest(eui, client.Keypair))
				if err != nil {
					panic(err)
				}
			}
			_, err = u.CloseAndRecv()
			if err != nil {
				panic(err)
			}
		}

		if len(skfs) > 0 {
			//TODO: limit to max 100 updates in a single call
			_, err := routeClient.UpdateSkfs(ctx, helium_crypto.SignRequest(&iot_config.RouteSkfUpdateReqV1{RouteId: routeId, Updates: skfs}, client.Keypair))
			if err != nil {
				panic(err)
			}
		}

		fmt.Printf("[helium] done transaction\n")
	}
}
