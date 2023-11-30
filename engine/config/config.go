package config

import (
	"os"

	"github.com/imdario/mergo"
)

var DefaultConfig = Config{
	SnapRestSeconds: 15,
}

func init() {
	wd, _ := os.Getwd()
	DefaultConfig.LogDir = wd
	DefaultConfig.SaveDir = wd
}

// Config 全局配置
type Config struct {
	LogDir          string
	SnapRestSeconds uint

	// 直播间配置的初始值
	SaveDir string `toml:"-"`
}

// CheckAndFix 检查配置，并合并默认配置
func (cfg *Config) CheckAndFix() {
	mergo.Merge(cfg, DefaultConfig)
}
