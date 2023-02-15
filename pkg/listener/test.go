package listener

import (
	"encoding/binary"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type TestListener struct {
	BaseListener
}

func (t *TestListener) Listen() {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()
	for _ = range ticker.C {
		id, _ := uuid.New().MarshalBinary()
		event := DeviceEvent{
			DevEui:    binary.BigEndian.Uint64(id[0:8]),
			AppEui:    binary.BigEndian.Uint64(id[8:16]),
			EventType: EventType_Added,
		}
		fmt.Printf("Generated: %+v\n", event)
		t.ch <- event
	}
}
