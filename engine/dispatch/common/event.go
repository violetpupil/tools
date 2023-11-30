package common

import "olive/engine/config"

type EventTypeID uint32

var EventType = struct {
	AddMonitor    EventTypeID
	RemoveMonitor EventTypeID

	AddRecorder    EventTypeID
	RemoveRecorder EventTypeID
}{
	AddMonitor:    101,
	RemoveMonitor: 102,

	AddRecorder:    201,
	RemoveRecorder: 202,
}

func (et EventTypeID) String() string {
	switch et {
	case EventType.AddMonitor:
		return "add monitor"
	case EventType.RemoveMonitor:
		return "remove monitor"
	case EventType.AddRecorder:
		return "add recorder"
	case EventType.RemoveRecorder:
		return "remove recorder"
	}
	return "undefined"
}

// Bout 直播间事件
type Bout interface {
	IsConfigValid() bool
	GetID() config.ID
	GetPlatform() string
	GetRoomID() string
}
