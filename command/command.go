package command

import "github.com/spf13/cobra"

// cmder 具体命令应该实现该接口，获取 cobra.Command 对象
type cmder interface {
	getCommand() *cobra.Command
}

// commandsBuilder 组合所有命令
type commandsBuilder struct {
	commands []cmder
}

func newCommandsBuilder() *commandsBuilder {
	return &commandsBuilder{}
}

func (b *commandsBuilder) addCommands(commands ...cmder) *commandsBuilder {
	b.commands = append(b.commands, commands...)
	return b
}

func (b *commandsBuilder) addAll() *commandsBuilder {
	// TODO
	return b
}

func Execute([]string) {}
