package kernel

type bout struct{}

func NewBout() *bout {
	return &bout{}
}

func (b *bout) AddMonitor() {}
