package config

import (
	"os"

	"github.com/imdario/mergo"
)

var DefaultConfig = Config{}

func init() {
	wd, _ := os.Getwd()
	DefaultConfig.LogDir = wd
	DefaultConfig.SaveDir = wd
}

type Config struct {
	LogDir  string
	SaveDir string
}

// CheckAndFix 检查配置，并合并默认配置
func (cfg *Config) CheckAndFix() {
	mergo.Merge(cfg, DefaultConfig)
}
