package kernel

import (
	"github.com/sirupsen/logrus"
)

type Kernel struct {
	log *logrus.Logger
}

func New(log *logrus.Logger) *Kernel {
	return &Kernel{
		log: log,
	}
}

func (k *Kernel) Run() {}
