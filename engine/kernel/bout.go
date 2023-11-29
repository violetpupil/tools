package kernel

import (
	"fmt"
	"olive/engine/config"
	"olive/foundation/olivetv"
	"olive/foundation/sync"
)

type Bout interface{}

// bout 直播间操作
type bout struct {
	*olivetv.TV

	cfg     *config.Config
	showMap *sync.RWMap[string, *Show]
	showID  string
	show    *Show
}

func NewBout(showID string, showMap *sync.RWMap[string, *Show], cfg *config.Config) (*bout, error) {
	show, ok := showMap.Get(showID)
	if !ok {
		return nil, fmt.Errorf("show[ID = %s] config does not exist", showID)
	}
	tv, err := olivetv.New(show.Platform, show.RoomID)
	if err != nil {
		return nil, err
	}

	return &bout{
		showID:  showID,
		showMap: showMap,
		cfg:     cfg,
		show:    show,
		TV:      tv,
	}, nil
}

func (b *bout) Refresh() {
	s, ok := b.showMap.Get(b.showID)
	if !ok {
		return
	}

	if s.Platform != b.SiteID || s.RoomID != b.RoomID {
		tv, err := olivetv.New(s.Platform, s.RoomID)
		if err != nil {
			return
		}
		b.TV = tv
	}

	b.show = s
}

func (b *bout) GetPlatform() string {
	b.Refresh()
	return b.SiteID
}

func (b *bout) GetRoomID() string {
	b.Refresh()
	return b.RoomID
}
