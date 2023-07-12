package lns

import (
	"github.com/thisisdevelopment/helium-route-updater/pkg/types"
)

type Client interface {
	GetDevices() []*types.Device
	Listen(ch chan<- types.DeviceEvent)
}

type BaseClient struct {
	c *types.LnsConfig
}

func NewClient(c *types.LnsConfig) Client {
	base := BaseClient{c: c}
	if c.Type == "chirpstack" {
		return NewChirpstackClient(base)
	} else {
		panic("Unsupported lns type: " + c.Type)
	}
}
