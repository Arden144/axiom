package main

import (
	"os"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/buttons"
	"github.com/arden144/axiom/commands"
	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"github.com/arden144/axiom/util"
	"github.com/arden144/axiom/utility"
)

func main() {
	defer log.L.Sync()
	defer bot.Client.Close(bot.Ctx)

	// bot.ClearCommands()
	util.AddCommands(config.DevGuildID, commands.Refresh, commands.Play, commands.PlayLink, commands.Skip, commands.Disconnect, commands.Pause, commands.Resume)
	bot.AddButtons(buttons.Toggle)

	utility.OnSignal(os.Interrupt)
}
