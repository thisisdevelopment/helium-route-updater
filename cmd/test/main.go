package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/thisisdevelopment/helium-route-updater/pkg/api/helium/service/iot_config"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"strings"
)

func main() {
	//input := "13WvV82S7QN3VMzMSieiGxvuaPKknMtf213E5JwPnboDkUfesKw"
	//keypair := helium_crypto.KeyPairFromString(input)
	//
	//signature, _ := hex.DecodeString("ef3e85dc7ea338c6b67399873131ea7b2265c516222e105fc39a59dda71f668a3b95fe27457d941a3cf5c422c9efbf0da112171d2997d74bc68f7b8118c6930e")
	//message := "hello world"
	//if keypair.Verify([]byte(message), signature) {
	//	fmt.Printf("This seems to work!\n")
	//}
	//
	//kp := helium_crypto.KeyPairGenerate(helium_crypto.NETTYPE_MAIN, helium_crypto.KEYTYPE_ED25519)
	//fmt.Printf("Public key: %s\n", kp.Public().String())
	//fmt.Printf("Private key: %s\n", hex.EncodeToString(kp.Bytes()))

	data, _ := os.ReadFile("./keypair.bin")
	keypair := helium_crypto.KeyPairFromBytes(data)
	//keypair := helium_crypto.KeyPairFromString("1pgvqN2QH52v4VtJZhkE4V24qzpqLX41soypvws4jDt3cKcEKg2zWK9U7MCwppEK4Jxfx1s9ZbcMK65SJaxYSBfVR7a8h4")
	fmt.Printf("Keypair: %s\n", keypair.String())
	fmt.Printf("Public: %s\n", keypair.Public().String())

	//server := "http://mainnet-config.helium.io:6080"
	server := "https://config.iot.mainnet.helium.io:6080"
	serverCredentials := credentials.NewTLS(&tls.Config{})
	if strings.HasPrefix(server, "http://") {
		server = server[7:]
		serverCredentials = insecure.NewCredentials()
	} else if strings.HasPrefix(server, "https://") {
		server = server[8:]
	}

	ctx := context.Background()
	conn, err := grpc.Dial(
		server,
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(serverCredentials),
	)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	//client := iot_config.NewOrgClient(conn)
	//res, err := client.List(ctx, &iot_config.OrgListReqV1{})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Printf("Result: %+v\n", res)

	rclient := iot_config.NewRouteClient(conn)
	rres, err := rclient.Get(ctx, helium_crypto.SignRequest(&iot_config.RouteGetReqV1{
		Id: "8344fd9c-09c2-11ee-a778-4b9462f9aeb4",
	}, keypair))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result: %+v\n", rres)
}
