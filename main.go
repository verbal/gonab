package main

import (
	"os"

	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/hobeone/gonab/commands"
)

func main() {
	commands.SetupCommands()
	kingpin.MustParse(commands.App.Parse(os.Args[1:]))
}
