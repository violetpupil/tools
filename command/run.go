package command

import (
	"olive/engine/config"
	"olive/engine/kernel"

	"github.com/spf13/cobra"
)

type runCmd struct {
	*baseCmd

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

	cmd.Flags().StringVarP(&cc.roomURL, "url", "u", "", "room url")
	cmd.Flags().StringVarP(&cc.proxy, "proxy", "p", "", "proxy url")

	return cc
}

func (c *runCmd) run() error {
	// TODO
	return nil
}

// CompositeConfig 配置文件
// 包括全局配置和每个直播间的配置
type CompositeConfig struct {
	Config config.Config
	Shows  []kernel.Show
}
