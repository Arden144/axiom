package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Disconnect = bot.Command{
	Create: bot.SlashCommand{
		Name:        "disconnect",
		Description: "disconnect",
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := bot.GetPlayer(*e.GuildID())

		if !player.State().Connected {
			msg.SetContent("not connected")
			return nil
		}

		if err := bot.Client.UpdateVoiceState(ctx, *e.GuildID(), nil, false, false); err != nil {
			return fmt.Errorf("failed to disconnect: %w", err)
		}
		player.Clear()

		msg.SetContent("disconnected")
		return nil
	},
}
