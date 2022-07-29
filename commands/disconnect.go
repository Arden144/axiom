package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var Disconnect = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "disconnect",
		Description: "disconnect",
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := music.Player(*e.GuildID())

		if !player.Connected() {
			msg.SetContent("not connected")
			return nil
		}

		if err := bot.Client.Disconnect(ctx, *e.GuildID()); err != nil {
			return fmt.Errorf("failed to disconnect: %w", err)
		}
		player.Clear()

		msg.SetContent("disconnected")
		return nil
	},
}
