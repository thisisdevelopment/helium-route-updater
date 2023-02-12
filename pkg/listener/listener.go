package listener

import "strings"

type EventType int64

func (e EventType) String() string {
	switch e {
	case EventType_Added:
		return "added"
	case EventType_Deleted:
		return "deleted"
	case EventType_Updated:
		return "updated"
	default:
		return "undefined"
	}
}

const (
	EventType_Undefined EventType = iota
	EventType_Added
	EventType_Updated
	EventType_Deleted
)

type DeviceEvent struct {
	DevEui    uint64
	AppEui    uint64
	EventType EventType
}

type Listener interface {
	Listen()
}

type BaseListener struct {
	ch     chan<- DeviceEvent
	listen string
}

func Listen(listen string, ch chan<- DeviceEvent) {
	var listener Listener
	base := BaseListener{listen: listen, ch: ch}
	if strings.HasPrefix(listen, "test://") {
		listener = &TestListener{BaseListener: base}
	} else {
		panic("unknown listener: " + listen)
	}
	listener.Listen()
}
