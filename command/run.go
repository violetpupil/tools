package command

import (
	"fmt"
	"olive/engine/config"
	"olive/engine/kernel"
	"olive/engine/log"
	"os"
	"time"

	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type runCmd struct {
	*baseCmd

	cfgFilepath string

	roomURL string
	proxy   string
}

func newRunCmd() *runCmd {
	cc := &runCmd{}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start the olive engine.",
		Long:  `Start the olive engine.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cc.run()
		},
	}
	cc.baseCmd = newBaseCmd(cmd)

	cmd.Flags().StringVarP(&cc.cfgFilepath, "filepath", "f", "", "set config.toml filepath")

	cmd.Flags().StringVarP(&cc.roomURL, "url", "u", "", "room url")
	cmd.Flags().StringVarP(&cc.proxy, "proxy", "p", "", "proxy url")

	return cc
}

func (c *runCmd) run() error {
	var cfg *CompositeConfig
	var err error
	if c.roomURL != "" {
		cfg, err = newCompositeConfigFromTerm(c)
	} else {
		cfg, err = newCompositeConfigFromFile(c)
	}
	if err != nil {
		return err
	}

	log.InitLogger(cfg.Config.LogDir)

	// TODO
	return nil
}

// CompositeConfig 配置文件
// 包括全局配置和每个直播间的配置
type CompositeConfig struct {
	Config config.Config
	Shows  []*kernel.Show
}

// newCompositeConfigFromTerm 从终端获取配置参数
func newCompositeConfigFromTerm(cmd *runCmd) (*CompositeConfig, error) {
	show, err := kernel.NewShow(cmd.roomURL, cmd.proxy)
	if err != nil {
		return nil, fmt.Errorf("invalid args: %w", err)
	}
	shows := []*kernel.Show{show}

	cfg := &CompositeConfig{
		Config: config.DefaultConfig,
		Shows:  shows,
	}
	cfg.checkAndFix()
	cfg.autosave()
	return cfg, nil
}

// newCompositeConfigFromFile 加载配置文件
func newCompositeConfigFromFile(cmd *runCmd) (*CompositeConfig, error) {
	viper.SetConfigFile(cmd.cfgFilepath)
	cfg := new(CompositeConfig)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	cfg.checkAndFix()
	return cfg, nil
}

// checkAndFix 合并默认全局配置
// 检查直播间字段，没有设置的话，设置初始值
func (cfg *CompositeConfig) checkAndFix() {
	cfg.Config.CheckAndFix()
	for _, show := range cfg.Shows {
		show.CheckAndFix(&cfg.Config)
	}
}

// autosave 保存配置文件
func (cfg *CompositeConfig) autosave() error {
	bs, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}
	err = os.WriteFile(fmt.Sprintf("config-%d.toml", time.Now().Unix()), bs, 0666)
	return err
}
