package lns

//
//import (
//	"encoding/binary"
//	"fmt"
//	"github.com/google/uuid"
//	"github.com/thisisdevelopment/helium-route-updater/pkg/listener"
//	"time"
//)
//
//type TestListener struct {
//	listener.BaseListener
//}
//
//func (t *TestListener) Listen() {
//	ticker := time.NewTicker(10 * time.Second)
//	defer ticker.Stop()
//	for _ = range ticker.C {
//		id, _ := uuid.New().MarshalBinary()
//		event := listener.DeviceEvent{
//			DevEui:    binary.BigEndian.Uint64(id[0:8]),
//			AppEui:    binary.BigEndian.Uint64(id[8:16]),
//			EventType: listener.EventType_Added,
//		}
//		fmt.Printf("Generated: %+v\n", event)
//		t.ch <- event
//	}
//}
