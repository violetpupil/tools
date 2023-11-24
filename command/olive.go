package command

import "github.com/spf13/cobra"

// 根命令
type oliveCmd struct {
	*baseCmd
}

func newOliveCmd() *oliveCmd {
	cc := &oliveCmd{}
	cc.baseCmd = newBaseCmd(&cobra.Command{
		Use:   "olive",
		Short: "olive is a live stream recorder",
		Long: `olive is a live stream recorder, underneath there is a powerful engine
which monitors streamers status and automatically records when they're 
online. It helps you catch every live stream you want to see.`,
	})
	return cc
}
