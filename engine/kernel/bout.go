package kernel

import (
	"fmt"
	"olive/engine/config"
	"olive/engine/dispatch/common"
	"olive/engine/dispatch/dispatcher"
	"olive/foundation/olivetv"
	"olive/foundation/sync"
)

// bout 直播间操作
type bout struct {
	*olivetv.TV

	cfg     *config.Config
	showMap *sync.RWMap[config.ID, *config.Show]
	showID  config.ID
	show    *config.Show
}

func NewBout(showID config.ID, showMap *sync.RWMap[config.ID, *config.Show], cfg *config.Config) (*bout, error) {
	show, ok := showMap.Get(showID)
	if !ok {
		return nil, fmt.Errorf("show[ID = %s] config does not exist", showID)
	}
	tv, err := olivetv.New(show.Platform, show.RoomID, olivetv.SetProxy(show.Proxy))
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
		tv, err := olivetv.New(s.Platform, s.RoomID, olivetv.SetProxy(s.Proxy))
		if err != nil {
			return
		}
		b.TV = tv
	}

	b.show = s
}

func (b *bout) IsConfigValid() bool {
	_, ok := b.showMap.Get(b.showID)
	return ok
}

func (b *bout) GetID() config.ID {
	return b.showID
}

func (b *bout) GetPlatform() string {
	b.Refresh()
	return b.SiteID
}

func (b *bout) GetRoomID() string {
	b.Refresh()
	return b.RoomID
}

// Snap 抓取直播间信息
func (b *bout) Snap() error {
	b.Refresh()
	return b.TV.Snap()
}

func (b *bout) RemoveMonitor() error {
	e := dispatcher.NewEvent(common.EventType.RemoveMonitor, b)
	return dispatcher.SharedManager.Dispatch(e)
}

func (b *bout) RemoveRecorder() error {
	e := dispatcher.NewEvent(common.EventType.RemoveRecorder, b)
	return dispatcher.SharedManager.Dispatch(e)
}
