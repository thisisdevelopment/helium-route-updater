package types

type Device struct {
	DevEui     uint64
	JoinEui    uint64
	DevAddr    uint32
	SessionKey []byte
}

type DeviceEvent struct {
	Update []*Device
	Delete []*Device
}
