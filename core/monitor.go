package core

import "github.com/go-olive/olive/foundation/olivetv"

type Monitor struct {
	site olivetv.Site
	tv   *olivetv.TV
}

func NewMonitor(site olivetv.Site, tv *olivetv.TV) *Monitor {
	return &Monitor{
		site: site,
		tv:   tv,
	}
}

func (m *Monitor) refresh() {}
