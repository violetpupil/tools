package monitor

import (
	"olive/engine/config"
	"olive/engine/dispatch/common"
	"olive/engine/dispatch/dispatcher"

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

func (d *Dispatcher) DispatcherType() common.DispatcherTypeID {
	return common.DispatcherType.Monitor
}

func (d *Dispatcher) DispatchTypes() []common.EventTypeID {
	return []common.EventTypeID{
		common.EventType.AddMonitor,
		common.EventType.RemoveMonitor,
	}
}
