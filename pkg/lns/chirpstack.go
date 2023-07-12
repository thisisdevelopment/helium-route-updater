package lns

import "github.com/thisisdevelopment/helium-route-updater/pkg/types"

type ChirpstackClient struct {
	BaseClient
}

func (c *ChirpstackClient) Listen(ch chan<- types.DeviceEvent) {

}

func (c *ChirpstackClient) GetDevices() []*types.Device {

}

func NewChirpstackClient(client BaseClient) *ChirpstackClient {
	return &ChirpstackClient{client}
}
