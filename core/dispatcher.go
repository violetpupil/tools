package core

import "github.com/go-olive/olive/foundation/olivetv"

type Dispatcher struct {
	site olivetv.Site
	tv   *olivetv.TV

	monitor  *monitor
	recorder *recorder
}

func NewDispatcher(site olivetv.Site, tv *olivetv.TV) *Dispatcher {
	return &Dispatcher{
		site:     site,
		tv:       tv,
		monitor:  newMonitor(site, tv),
		recorder: newRecorder(site, tv),
	}
}
