package main

import (
	"context"
	"os"
	"strings"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/buttons"
	"github.com/arden144/axiom/commands"
	"github.com/arden144/axiom/log"
	"github.com/arden144/axiom/search"
	"github.com/arden144/axiom/slices"
	"github.com/arden144/axiom/utility"
	"go.uber.org/zap"
)

func main() {
	defer log.L.Sync()
	defer bot.Client.Close(context.Background())

	track, err := search.Search("jimmy cooks")
	if err != nil {
		log.L.Fatal("search failed", zap.Error(err))
	}

	name := track.Name
	artistNames := slices.Map(track.Artists, func(a search.Artist) string { return a.Name })
	artists := strings.Join(artistNames, ", ")

	log.S.Infof("%v - %v", name, artists)

	// bot.ClearCommands()
	bot.AddCommands(commands.Play, commands.Skip, commands.Disconnect, commands.Pause, commands.Resume)
	bot.AddButtons(buttons.Pause)

	utility.OnSignal(os.Interrupt)
}
