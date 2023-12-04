package command

import "github.com/spf13/cobra"

// run run 子命令
type run struct {
	roomURL string
}

func newRun() *cobra.Command {
	c := new(run)

	cc := &cobra.Command{
		Use:   "run",
		Short: "Start the olive engine.",
		Long:  `Start the olive engine.`,
		Run:   c.run,
	}

	cc.Flags().StringVarP(&c.roomURL, "url", "u", "", "room url")

	return cc
}

// run 执行函数
func (c *run) run(*cobra.Command, []string) {
	// TODO
}
