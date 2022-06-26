package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Pause = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "pause",
		Description: "pause",
	},
	Handler: func(_ context.Context, b *bot.Bot, e bot.CommandEvent) (*discord.MessageUpdate, error) {
		player := b.Music.Player(*e.GuildID())

		if !player.Playing() {
			return e.Reply("nothing to pause")
		}

		if player.Paused() {
			return e.Reply("already paused")
		}

		if err := player.Pause(true); err != nil {
			return e.Fatal("failed to pause", err)
		}

		return e.Reply("paused")
	},
}
