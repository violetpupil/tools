package monitor

import (
	"olive/engine/config"
	"olive/engine/dispatch/common"
	"olive/engine/dispatch/dispatcher"
	"sync/atomic"
	"time"

	"github.com/lthibault/jitterbug"
	"github.com/sirupsen/logrus"
)

type monitor struct {
	status common.StatusID
	bout   common.Bout
	// 轮询时通过 stop 停止监控
	// 停止后通过 done 表明结束
	stop chan struct{}
	done chan struct{}

	log *logrus.Logger
	cfg *config.Config

	roomOn bool
}

func newMonitor(log *logrus.Logger, bout common.Bout, cfg *config.Config) *monitor {
	return &monitor{
		status: common.Status.Starting,
		bout:   bout,
		stop:   make(chan struct{}),
		done:   make(chan struct{}),

		log: log,
		cfg: cfg,
	}
}

func (m *monitor) Start() error {
	if !atomic.CompareAndSwapUint32(&m.status, common.Status.Starting, common.Status.Pending) {
		return nil
	}
	m.log.WithFields(logrus.Fields{
		"pf": m.bout.GetPlatform(),
		"id": m.bout.GetRoomID(),
	}).Info("monitor start")
	defer atomic.CompareAndSwapUint32(&m.status, common.Status.Pending, common.Status.Running)

	m.refresh()
	go m.run()
	return nil
}

func (m *monitor) Stop() {
	if !atomic.CompareAndSwapUint32(&m.status, common.Status.Running, common.Status.Stopping) {
		return
	}
	close(m.stop)
}

func (m *monitor) refresh() {
	// 直播间被移除
	if !m.bout.IsConfigValid() {
		m.Stop()
		return
	}

	if err := m.bout.Snap(); err != nil {
		m.log.WithFields(logrus.Fields{
			"pf": m.bout.GetPlatform(),
			"id": m.bout.GetRoomID(),
		}).Tracef("snap failed, %s", err)
		return
	}
	_, roomOn := m.bout.StreamURL()
	defer func() { m.roomOn = roomOn }()
	if m.roomOn || !roomOn {
		return
	}

	m.log.WithFields(logrus.Fields{
		"pf":  m.bout.GetPlatform(),
		"id":  m.bout.GetRoomID(),
		"old": m.roomOn,
		"new": roomOn,
	}).Info("live status changed")

	e := dispatcher.NewEvent(common.EventType.AddRecorder, m.bout)
	if err := dispatcher.SharedManager.Dispatch(e); err != nil {
		m.log.Error(err)
	}
}

func (m *monitor) run() {
	t := jitterbug.New(
		time.Duration(m.cfg.SnapRestSeconds)*time.Second,
		&jitterbug.Norm{Stdev: 3 * time.Second},
	)
	defer t.Stop()

	for {
		select {
		case <-m.stop:
			m.log.WithFields(logrus.Fields{
				"pf": m.bout.GetPlatform(),
				"id": m.bout.GetRoomID(),
			}).Info("monitor stop")
			close(m.done)
			return
		case <-t.C:
			m.refresh()
		}
	}
}
