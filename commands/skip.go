package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Skip = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "skip",
		Description: "skip",
	},
	Handler: func(ctx context.Context, b *bot.Bot, e bot.CommandEvent) (*discord.MessageUpdate, error) {
		player := b.Music.Player(*e.GuildID())

		if !player.Playing() {
			return e.Reply("nothing to skip")
		}

		if err := player.Stop(); err != nil {
			return e.Fatal("failed to stop", err)
		}

		if err := player.Next(); err != nil {
			return e.Fatal("failed to play", err)
		}

		return e.Reply("skipped")
	},
}
