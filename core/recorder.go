package core

import (
	"github.com/go-olive/olive/foundation/olivetv"
	"github.com/sirupsen/logrus"
)

type Recorder struct {
	site olivetv.Site
	tv   *olivetv.TV

	log *logrus.Entry
}

func NewRecorder(site olivetv.Site, tv *olivetv.TV) *Recorder {
	return &Recorder{
		site: site,
		tv:   tv,
		log: Logger.WithFields(logrus.Fields{
			"pf": tv.SiteID,
			"id": tv.RoomID,
		}),
	}
}
