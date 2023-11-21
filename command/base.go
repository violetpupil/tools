package command

import "github.com/spf13/cobra"

// baseCmd 封装 cobra.Command
// 统一实现 cobra.Command 获取
// 具体命令通过嵌入实现 cmder 接口
type baseCmd struct {
	cmd *cobra.Command
}

func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}
