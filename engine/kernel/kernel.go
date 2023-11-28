package kernel

type Kernel struct{}

func New() *Kernel {
	return &Kernel{}
}

func (k *Kernel) Run() {}
