package dbwriter

import (
	"github.com/patricknoir/F12020/pkg/common/packet"
	"sync"
)

type SessionData struct {
	lapDistance float32
	mutex sync.Mutex
}

func Write(data packet.LapDataPacket) error {
	return nil
}