package main

import (
	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/buttons"
	"github.com/arden144/axiom/commands"
)

func main() {
	// bot.ClearCommands()
	bot.AddCommands(commands.Play, commands.Skip, commands.Disconnect, commands.Pause, commands.Resume)
	bot.AddButtons(buttons.Pause)
}
