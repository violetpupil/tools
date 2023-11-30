package monitor

import (
	"errors"
	"olive/engine/config"
	"olive/engine/dispatch/common"
	"olive/engine/dispatch/dispatcher"
	"sync"

	"github.com/sirupsen/logrus"
)

type Dispatcher struct {
	log *logrus.Logger
	cfg *config.Config

	mu     sync.RWMutex
	savers map[config.ID]*monitor
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
	bout.RemoveRecorder()

	d.mu.Lock()
	defer d.mu.Unlock()

	if _, ok := d.savers[bout.GetID()]; ok {
		return errors.New("exist")
	}
	monitor := newMonitor(d.log, bout, d.cfg)
	d.savers[bout.GetID()] = monitor
	return monitor.Start()
}

func (d *Dispatcher) removeMonitor(bout common.Bout) error {
	d.mu.Lock()
	defer d.mu.Unlock()

	monitor, ok := d.savers[bout.GetID()]
	if !ok {
		return errors.New("monitor not exist")
	}
	monitor.Stop()
	delete(d.savers, bout.GetID())
	return nil
}
