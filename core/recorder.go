package core

import (
	"github.com/go-olive/olive/foundation/olivetv"
	"github.com/sirupsen/logrus"
)

type recorder struct {
	site olivetv.Site
	tv   *olivetv.TV

	log *logrus.Entry
}

func newRecorder(site olivetv.Site, tv *olivetv.TV) *recorder {
	return &recorder{
		site: site,
		tv:   tv,
		log: logger.WithFields(logrus.Fields{
			"pf": tv.SiteID,
			"id": tv.RoomID,
		}),
	}
}
