package main

import (
	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
	"github.com/thisisdevelopment/helium-route-updater/pkg/updater"
)

func main() {
	//TODO: from config
	routeId := "test"
	listen := "test://" + routeId
	server := "test-server:8000"
	ch := make(chan listener.DeviceEvent)
	go listener.Listen(listen, ch)
	updater.Run(server, routeId, ch)
}
