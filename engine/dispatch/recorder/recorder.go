package recorder

import (
	"olive/engine/dispatch/common"

	"github.com/sirupsen/logrus"
)

type recorder struct {
	bout   common.Bout
	log    *logrus.Logger
	status common.StatusID
	stop   chan struct{}
	done   chan struct{}
}

func newRecorder(log *logrus.Logger, bout common.Bout) (*recorder, error) {
	return &recorder{
		status: common.Status.Starting,
		bout:   bout,
		log:    log,
		stop:   make(chan struct{}),
		done:   make(chan struct{}),
	}, nil
}

func (r *recorder) record() error {
	// TODO
	return nil
}
