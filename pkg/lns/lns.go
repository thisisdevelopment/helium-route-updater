package lns

import (
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

type Client interface {
	GetDevices() []*types.Device
	GetDevice(deviceId string) (*types.Device, error)
	Listen(ch chan<- types.DeviceEvent)
}

type BaseClient struct {
	config *types.LnsConfig
}

func NewClient(c *types.LnsConfig) Client {
	base := BaseClient{config: c}
	if c.Type == "chirpstack" {
		return NewChirpstackClient(base)
	} else {
		panic("Unsupported lns type: " + c.Type)
	}
}
