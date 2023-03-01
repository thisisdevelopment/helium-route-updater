package updater

import (
	"context"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func Run(server string, routeId string, keypair *helium_crypto.KeyPair, ch <-chan listener.DeviceEvent) {
	ctx := context.Background()
	conn, err := grpc.Dial(
		server,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := iot_config.NewRouteClient(conn)

	for event := range ch {
		pair := &iot_config.EuiPairV1{
			RouteId: routeId,
			AppEui:  event.AppEui,
			DevEui:  event.DevEui,
		}

		u, err := client.UpdateEuis(ctx)
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
			}, keypair))
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
			}, keypair))
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
