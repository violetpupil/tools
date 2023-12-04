package command

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var root = &cobra.Command{
	Use:   "olive",
	Short: "olive is a live stream recorder",
	Long: `olive is a live stream recorder, underneath there is a powerful engine
which monitors streamers status and automatically records when they're 
online. It helps you catch every live stream you want to see.`,
}

func Execute() {
	root.AddCommand(newRun())
	err := root.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
