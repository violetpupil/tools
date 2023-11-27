package log

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func InitLogger(logDir string) {
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatal(err)
	}
	// 追加日志
	f, err := os.OpenFile(
		filepath.Join(logDir, "l.log"),
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0660,
	)
	if err != nil {
		log.Fatal(err)
	}

	// 终端输出颜色
	Logger := &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.TextFormatter{
			ForceColors:     true,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		},
		Hooks: make(logrus.LevelHooks),
		Level: logrus.TraceLevel,
	}

	// 文件只输出文本
	Logger.AddHook(lfshook.NewHook(
		f,
		&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		},
	))
}
