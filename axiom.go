package main

import (
	"os"

	_ "github.com/arden144/axiom/web"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/buttons"
	"github.com/arden144/axiom/commands"
	"github.com/arden144/axiom/log"
	"github.com/arden144/axiom/utility"
)

func main() {
	defer log.L.Sync()
	defer bot.Client.Close(bot.Ctx)

	// bot.ClearCommands()
	bot.AddCommands(commands.Play, commands.Skip, commands.Disconnect, commands.Pause, commands.Resume)
	bot.AddButtons(buttons.Pause)

	utility.OnSignal(os.Interrupt)
}
