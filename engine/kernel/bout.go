package kernel

import "olive/foundation/olivetv"

type Bout interface{}

// bout 直播间操作
type bout struct {
	*olivetv.TV
}

func NewBout() *bout {
	return &bout{}
}

func (b *bout) AddMonitor() {}
