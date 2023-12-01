package recorder

import (
	"errors"
	"fmt"
	"olive/engine/dispatch/common"
	"olive/engine/parser"
	"sync/atomic"
	"time"

	"github.com/sirupsen/logrus"
)

type recorder struct {
	bout   common.Bout
	log    *logrus.Logger
	status common.StatusID
	stop   chan struct{}
	done   chan struct{}
	parser parser.Parser
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

func (r *recorder) Start() error {
	if !atomic.CompareAndSwapUint32(&r.status, common.Status.Starting, common.Status.Pending) {
		return nil
	}
	defer atomic.CompareAndSwapUint32(&r.status, common.Status.Pending, common.Status.Running)
	go r.run()

	r.log.WithFields(logrus.Fields{
		"pf": r.bout.GetPlatform(),
		"id": r.bout.GetRoomID(),
	}).Info("recorder start")

	return nil
}

func (r *recorder) Stop() {
	if !atomic.CompareAndSwapUint32(&r.status, common.Status.Running, common.Status.Stopping) {
		return
	}
	close(r.stop)
	if r.parser != nil {
		r.parser.Stop()
	}
}

func (r *recorder) record() error {
	if !r.bout.IsConfigValid() {
		r.Stop()
		return nil
	}

	newParser, exist := parser.SharedManager.Parser(r.bout.GetParser())
	if !exist {
		return fmt.Errorf("parser[%s] does not exist", r.bout.GetParser())
	}
	r.parser = newParser.New()

	// TODO

	const retry = 3
	var streamURL string
	var ok bool
	for i := 0; i < retry; i++ {
		err := r.bout.Snap()
		if err == nil {
			if streamURL, ok = r.bout.StreamURL(); ok {
				break
			} else {
				err = errors.New("empty stream url")
			}
			r.log.WithFields(logrus.Fields{
				"pf":  r.bout.GetPlatform(),
				"id":  r.bout.GetRoomID(),
				"cnt": i + 1,
			}).Errorf("snap failed, %s", err)

			if i == retry-1 {
				return err
			}
			time.Sleep(5 * time.Second)
		}
	}

	// TODO
	fmt.Println(streamURL)
	return nil
}

func (r *recorder) run() {
	r.bout.RemoveMonitor()

	defer func() {
		select {
		case <-r.stop:
		default:
			r.bout.AddMonitor()
		}
	}()

	for {
		select {
		case <-r.stop:
			close(r.done)
			r.log.WithFields(logrus.Fields{
				"pf": r.bout.GetPlatform(),
				"id": r.bout.GetRoomID(),
			}).Info("recorder stop")
			return
		default:
			if err := r.record(); err != nil {
				return
			}
		}
	}
}
