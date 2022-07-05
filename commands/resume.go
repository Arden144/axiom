package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Resume = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "resume",
		Description: "resume",
	},
	Handler: func(_ context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := e.Bot.Music.Player(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("nothing to resume")
			return nil
		}

		if !player.Paused() {
			msg.SetContent("already playing")
			return nil
		}

		if err := player.Pause(false); err != nil {
			return fmt.Errorf("failed to resume: %w", err)
		}

		msg.SetContent("resumed")
		return nil
	},
}
