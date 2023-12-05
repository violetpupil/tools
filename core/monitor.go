package core

import (
	"github.com/go-olive/olive/foundation/olivetv"
	"github.com/sirupsen/logrus"
)

type Monitor struct {
	site olivetv.Site
	tv   *olivetv.TV

	log *logrus.Entry
}

func NewMonitor(site olivetv.Site, tv *olivetv.TV) *Monitor {
	return &Monitor{
		site: site,
		tv:   tv,
		log: Logger.WithFields(logrus.Fields{
			"pf": tv.SiteID,
			"id": tv.RoomID,
		}),
	}
}

func (m *Monitor) refresh() {
	err := m.site.Snap(m.tv)
	if err != nil {
		m.log.Tracef("snap failed, %s", err)
		return
	}
	_, roomOn := m.tv.StreamURL()
	if !roomOn {
		return
	}

	// TODO
}
