package main

import (
	"encoding/hex"
	helium_crypto "github.com/thisisdevelopment/helium-route-updater/pkg/helium-crypto"
	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
	"github.com/thisisdevelopment/helium-route-updater/pkg/updater"
)

func main() {
	//TODO: from config
	bytes, _ := hex.DecodeString("012033ec6b9a163247609f89f91599166e1fb46f342f8fda33667db65a606f80af71a8aa142eaa0c3889e8cbc2d00a1a102cc2d7a6977fa9a34f0bedbb9635c7f5")
	keypair := helium_crypto.KeyPairFromBytes(bytes)

	routeId := "test"
	listen := "test://" + routeId
	server := "test-server:8000"
	ch := make(chan listener.DeviceEvent)
	go listener.Listen(listen, ch)
	updater.Run(server, routeId, keypair, ch)
}
