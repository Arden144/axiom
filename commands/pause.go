package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var Pause = bot.Command{
	Create: bot.SlashCommand{
		Name:        "pause",
		Description: "pause",
	},
	Handler: func(_ context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		player := music.GetPlayer(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("nothing to pause")
			return nil
		}

		if player.Paused() {
			msg.SetContent("already paused")
			return nil
		}

		if err := player.Pause(true); err != nil {
			return fmt.Errorf("failed to pause: %w", err)
		}

		msg.SetEmbeds(embeds.Pause(player.PlayingTrack().Info(), player.PlayingTrack().Info().Length-player.Position()))
		return nil
	},
}
