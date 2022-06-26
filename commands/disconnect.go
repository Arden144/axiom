package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Disconnect = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "disconnect",
		Description: "disconnect",
	},
	Handler: func(ctx context.Context, b *bot.Bot, e bot.CommandEvent) (*discord.MessageUpdate, error) {
		player := b.Music.Player(*e.GuildID())

		if !player.Connected() {
			return e.Reply("not connected")
		}

		if err := b.Client.Disconnect(ctx, *e.GuildID()); err != nil {
			return e.Fatal("failed to disconnect", err)
		}
		player.Clear()

		return e.Reply("disconnected")
	},
}
