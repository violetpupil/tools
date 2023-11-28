package dispatcher

import (
	"errors"
	"olive/engine/dispatch/enum"

	"github.com/sirupsen/logrus"
)

var SharedManager *Manager

type Manager struct {
	savers             map[enum.DispatcherTypeID]Dispatcher
	dispatchFuncSavers map[enum.EventTypeID]Dispatcher

	log *logrus.Logger
}

func NewManager(log *logrus.Logger) *Manager {
	return &Manager{
		savers:             make(map[enum.DispatcherTypeID]Dispatcher),
		dispatchFuncSavers: make(map[enum.EventTypeID]Dispatcher),
		log:                log,
	}
}

func (m *Manager) Register(ds ...Dispatcher) {
	if m.savers == nil {
		m.savers = map[enum.DispatcherTypeID]Dispatcher{}
	}
	if m.dispatchFuncSavers == nil {
		m.dispatchFuncSavers = map[enum.EventTypeID]Dispatcher{}
	}

	for _, d := range ds {
		t := d.DispatcherType()
		_, ok := m.savers[t]
		if ok {
			m.log.WithField("type", t).Warn("dispatcher has registered")
		}
		m.savers[t] = d

		for _, v := range d.DispatchTypes() {
			m.RegisterFunc(v, d)
		}
	}
}

func (m *Manager) RegisterFunc(typ enum.EventTypeID, d Dispatcher) {
	if m.dispatchFuncSavers == nil {
		m.dispatchFuncSavers = map[enum.EventTypeID]Dispatcher{}
	}
	_, ok := m.dispatchFuncSavers[typ]
	if ok {
		m.log.WithField("type", typ).Warn("dispatch func has registered")
	}
	m.dispatchFuncSavers[typ] = d
}

func (m *Manager) Dispatcher(typ enum.DispatcherTypeID) (Dispatcher, bool) {
	v, ok := m.savers[typ]
	return v, ok
}

func (m *Manager) Dispatch(e *Event) error {
	dispatcher, ok := m.dispatchFuncSavers[e.Type]
	if !ok {
		return errors.New("dispatch func does not exist")
	}
	return dispatcher.Dispatch(e)
}
