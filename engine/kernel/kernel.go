package kernel

import (
	"olive/engine/config"
	"olive/engine/dispatch/dispatcher"
	"olive/engine/dispatch/monitor"
	"olive/engine/dispatch/recorder"
	"olive/foundation/sync"

	"github.com/sirupsen/logrus"
)

type Kernel struct {
	log     *logrus.Logger
	cfg     *config.Config
	showMap *sync.RWMap[string, *Show]

	recorderDispatcher *recorder.Dispatcher
	monitorDispatcher  *monitor.Dispatcher
}

func New(log *logrus.Logger, cfg *config.Config, shows []*Show) *Kernel {
	showMap := sync.NewRWMap[string, *Show](len(shows))
	for _, show := range shows {
		showMap.Set(show.ID, show)
	}

	recorderDispatcher := recorder.New(log, cfg)
	monitorDispatcher := monitor.New(log, cfg)
	dispatcher.SharedManager = dispatcher.NewManager(log)
	dispatcher.SharedManager.Register(recorderDispatcher, monitorDispatcher)

	return &Kernel{
		log:     log,
		cfg:     cfg,
		showMap: showMap,

		recorderDispatcher: recorderDispatcher,
		monitorDispatcher:  monitorDispatcher,
	}
}

func (k *Kernel) UpdateConfig() {}

func (k *Kernel) UpdateShow() {}

func (k *Kernel) Run() {}
