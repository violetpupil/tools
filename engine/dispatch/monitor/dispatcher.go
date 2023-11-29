package monitor

import (
	"olive/engine/config"
	"olive/engine/dispatch/dispatcher"
	"olive/engine/dispatch/enum"

	"github.com/sirupsen/logrus"
)

type Dispatcher struct {
	log *logrus.Logger
	cfg *config.Config
}

func New(log *logrus.Logger, cfg *config.Config) *Dispatcher {
	return &Dispatcher{
		log: log,
		cfg: cfg,
	}
}

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
