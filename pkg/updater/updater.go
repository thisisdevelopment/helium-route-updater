package updater

import (
	"context"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_api "github.com/thisisdevelopment/helium-route-updater/pkg/helium-api"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
	"log"
)

func Run(client helium_api.Client, routeId string, ch <-chan listener.DeviceEvent) {
	ctx := context.Background()
	routeClient := client.NewRouteClient()
	for event := range ch {
		pair := &iot_config.EuiPairV1{
			RouteId: routeId,
			AppEui:  event.AppEui,
			DevEui:  event.DevEui,
		}

		u, err := routeClient.UpdateEuis(ctx)
		if err != nil {
			log.Fatal(err)
			return
		}

		if event.EventType == listener.EventType_Updated || event.EventType == listener.EventType_Deleted {
			err := u.Send(helium_crypto.SignRequest(&iot_config.RouteUpdateEuisReqV1{
				Action:    iot_config.ActionV1_remove,
				EuiPair:   pair,
				Timestamp: 0,
				Signature: nil,
			}, client.Keypair))
			if err != nil {
				log.Fatal(err)
				return
			}
		}
		if event.EventType == listener.EventType_Updated || event.EventType == listener.EventType_Added {
			err := u.Send(helium_crypto.SignRequest(&iot_config.RouteUpdateEuisReqV1{
				Action:    iot_config.ActionV1_add,
				EuiPair:   pair,
				Timestamp: 0,
				Signature: nil,
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
	}
}
