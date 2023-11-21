package command

import "github.com/spf13/cobra"

type runCmd struct {
	*baseCmd
}

func newRunCmd() *runCmd {
	cc := &runCmd{}

	// TODO
	cmd := &cobra.Command{
		Use:   "run",
		Short: "Start the olive engine.",
		Long:  `Start the olive engine.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cc.run()
		},
	}
	cc.baseCmd = newBaseCmd(cmd)

	return cc
}

func (c *runCmd) run() error {
	// TODO
	return nil
}
