package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Resume = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "resume",
		Description: "resume",
	},
	Handler: func(_ context.Context, b *bot.Bot, e bot.CommandEvent) (*discord.MessageUpdate, error) {
		player := b.Music.Player(*e.GuildID())

		if !player.Playing() {
			return e.Reply("nothing to resume")
		}

		if !player.Paused() {
			return e.Reply("already playing")
		}

		if err := player.Pause(false); err != nil {
			return e.Fatal("failed to resume", err)
		}

		return e.Reply("resumed")
	},
}
