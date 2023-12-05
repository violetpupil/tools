package core

import (
	"time"

	"github.com/go-olive/olive/foundation/olivetv"
	"github.com/lthibault/jitterbug"
	"github.com/sirupsen/logrus"
)

type monitor struct {
	site olivetv.Site
	tv   *olivetv.TV

	log  *logrus.Entry
	stop chan struct{}
}

func newMonitor(site olivetv.Site, tv *olivetv.TV) *monitor {
	return &monitor{
		site: site,
		tv:   tv,
		log: logger.WithFields(logrus.Fields{
			"pf": tv.SiteID,
			"id": tv.RoomID,
		}),
		stop: make(chan struct{}),
	}
}

func (m *monitor) Start() {
	m.log.Info("monitor start")
	m.refresh()
	m.run()
}

func (m *monitor) refresh() {
	err := m.site.Snap(m.tv)
	if err != nil {
		m.log.Tracef("snap failed, %s", err)
		return
	}
	_, roomOn := m.tv.StreamURL()
	if !roomOn {
		return
	}
	close(m.stop)
}

func (m *monitor) run() {
	t := jitterbug.New(
		15*time.Second,
		&jitterbug.Norm{Stdev: 3 * time.Second},
	)
	defer t.Stop()

	for {
		select {
		case <-m.stop:
			m.log.Info("monitor stop")
			break
		case <-t.C:
			m.refresh()
		}
	}
}
