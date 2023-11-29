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
	bout := event.Object.(common.Bout)

	d.log.WithFields(logrus.Fields{
		"pf": bout.GetPlatform(),
		"id": bout.GetRoomID(),
	}).Infoln("dispatch", event.Type)

	switch event.Type {
	case common.EventType.AddMonitor:
		return d.addMonitor(bout)
	case common.EventType.RemoveMonitor:
		return d.removeMonitor(bout)
	}
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

func (d *Dispatcher) addMonitor(bout common.Bout) error {
	// TODO
	return nil
}

func (d *Dispatcher) removeMonitor(bout common.Bout) error {
	// TODO
	return nil
}
