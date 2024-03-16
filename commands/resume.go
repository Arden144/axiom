package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var Resume = bot.Command{
	Create: bot.SlashCommand{
		Name:        "resume",
		Description: "resume",
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := bot.GetPlayer(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("nothing to resume")
			return nil
		}

		if !player.Paused() {
			msg.SetContent("already playing")
			return nil
		}

		if err := player.Update(ctx, lavalink.WithPaused(false)); err != nil {
			return fmt.Errorf("failed to resume: %w", err)
		}

		msg.SetContent("resumed")
		return nil
	},
}
