package main

import (
	"olive/command"
	"os"
)

func main() {
	command.Execute(os.Args[1:])
}
