package sites

import "sync"

type Site interface{}

var sites sync.Map

func registerSite(siteID string, site Site) {
	if _, dup := sites.LoadOrStore(siteID, site); dup {
		panic("site already registered")
	}
}

func Sniff(siteID string) (Site, bool) {
	s, ok := sites.Load(siteID)
	if !ok {
		return nil, false
	}
	return s.(Site), true
}
