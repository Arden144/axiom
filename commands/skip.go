package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var Skip = bot.Command{
	Create: bot.SlashCommand{
		Name:        "skip",
		Description: "skip",
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := music.GetPlayer(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("nothing to skip")
			return nil
		}

		if err := player.Stop(); err != nil {
			return fmt.Errorf("failed to stop: %w", err)
		}

		if err := player.Next(); err != nil {
			return fmt.Errorf("failed to play: %w", err)
		}

		msg.SetContent("skipped")
		return nil
	},
}
