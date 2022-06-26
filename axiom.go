package main

import (
	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/commands"
	"github.com/arden144/axiom/config"
)

func main() {
	config := config.Read("config.json")
	bot := bot.New(config)
	bot.Setup()

	// bot.ClearCommands()
	bot.AddCommand(commands.Play)
	bot.AddCommand(commands.Skip)
	bot.AddCommand(commands.Disconnect)
	bot.AddCommand(commands.Pause)
	bot.AddCommand(commands.Resume)

	bot.Start()
}
