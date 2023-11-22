package olivetv

import "sync"

var sites sync.Map

type Site interface {
	Name() string
	RoomID(RoomURL) string
}

func registerSite(siteID string, site Site) {
	if _, dup := sites.LoadOrStore(siteID, site); dup {
		panic("site already registered")
	}
}

// Sniff 获取网站实例
func Sniff(siteID string) (Site, bool) {
	s, ok := sites.Load(siteID)
	if !ok {
		return nil, ok
	}
	return s.(Site), ok
}
