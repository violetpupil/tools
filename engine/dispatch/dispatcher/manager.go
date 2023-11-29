package dispatcher

import (
	"errors"
	"olive/engine/dispatch/common"

	"github.com/sirupsen/logrus"
)

var SharedManager *Manager

type Manager struct {
	savers             map[common.DispatcherTypeID]Dispatcher
	dispatchFuncSavers map[common.EventTypeID]Dispatcher

	log *logrus.Logger
}

func NewManager(log *logrus.Logger) *Manager {
	return &Manager{
		savers:             make(map[common.DispatcherTypeID]Dispatcher),
		dispatchFuncSavers: make(map[common.EventTypeID]Dispatcher),
		log:                log,
	}
}

func (m *Manager) Register(ds ...Dispatcher) {
	if m.savers == nil {
		m.savers = map[common.DispatcherTypeID]Dispatcher{}
	}
	if m.dispatchFuncSavers == nil {
		m.dispatchFuncSavers = map[common.EventTypeID]Dispatcher{}
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

func (m *Manager) RegisterFunc(typ common.EventTypeID, d Dispatcher) {
	if m.dispatchFuncSavers == nil {
		m.dispatchFuncSavers = map[common.EventTypeID]Dispatcher{}
	}
	_, ok := m.dispatchFuncSavers[typ]
	if ok {
		m.log.WithField("type", typ).Warn("dispatch func has registered")
	}
	m.dispatchFuncSavers[typ] = d
}

func (m *Manager) Dispatcher(typ common.DispatcherTypeID) (Dispatcher, bool) {
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
