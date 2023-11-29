package monitor

import (
	"olive/engine/dispatch/dispatcher"
	"olive/engine/dispatch/enum"
)

type Dispatcher struct{}

func (d *Dispatcher) Dispatch(event *dispatcher.Event) error {
	// TODO
	return nil
}

func (d *Dispatcher) DispatcherType() enum.DispatcherTypeID {
	return enum.DispatcherType.Monitor
}

func (d *Dispatcher) DispatchTypes() []enum.EventTypeID {
	return []enum.EventTypeID{
		enum.EventType.AddMonitor,
		enum.EventType.RemoveMonitor,
	}
}
