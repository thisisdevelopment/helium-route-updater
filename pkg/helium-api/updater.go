package helium_api

import (
	"context"
	"encoding/hex"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
	"log"
)

func Run(client *Client, routeId string, ch <-chan types.DeviceEvent) {
	ctx := context.Background()
	routeClient := client.NewRouteClient()

	//TODO: better error handling
	for event := range ch {

		u, err := routeClient.UpdateEuis(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		skf := &iot_config.RouteSkfUpdateReqV1{RouteId: routeId, Updates: []*iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{}}
		for _, dev := range event.Update {
			if dev.DevAddr != 0 && len(dev.SessionKey) == 0 {
				skf.Updates = append(skf.Updates, &iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{
					Devaddr:    dev.DevAddr,
					SessionKey: hex.EncodeToString(dev.SessionKey),
					Action:     iot_config.ActionV1_add,
				})
			}
			if dev.JoinEui == 0 || dev.DevEui == 0 {
				continue
			}

			pair := &iot_config.EuiPairV1{
				RouteId: routeId,
				AppEui:  dev.JoinEui,
				DevEui:  dev.DevEui,
			}

			err := u.Send(helium_crypto.SignRequest(&iot_config.RouteUpdateEuisReqV1{
				Action:  iot_config.ActionV1_add,
				EuiPair: pair,
			}, client.Keypair))
			if err != nil {
				log.Fatal(err)
				return
			}
		}

		for _, dev := range event.Delete {
			if dev.DevAddr != 0 && len(dev.SessionKey) == 0 {
				skf.Updates = append(skf.Updates, &iot_config.RouteSkfUpdateReqV1RouteSkfUpdateV1{
					Devaddr:    dev.DevAddr,
					SessionKey: hex.EncodeToString(dev.SessionKey),
					Action:     iot_config.ActionV1_remove,
				})
			}
			if dev.JoinEui == 0 || dev.DevEui == 0 {
				continue
			}

			pair := &iot_config.EuiPairV1{
				RouteId: routeId,
				AppEui:  dev.JoinEui,
				DevEui:  dev.DevEui,
			}

			err := u.Send(helium_crypto.SignRequest(&iot_config.RouteUpdateEuisReqV1{
				Action:  iot_config.ActionV1_remove,
				EuiPair: pair,
			}, client.Keypair))
			if err != nil {
				log.Fatal(err)
				return
			}
		}

		_, err = u.CloseAndRecv()
		if err != nil {
			log.Fatal(err)
			return
		}

		//TODO: limit to max 100 updates in a single call
		_, err = routeClient.UpdateSkfs(ctx, helium_crypto.SignRequest(skf, client.Keypair))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
