package buttons

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var Toggle = bot.Button{
	Query: "toggle",
	Handler: func(_ context.Context, e bot.ButtonEvent, msg *discord.MessageCreateBuilder) error {
		player := music.GetPlayer(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("There's nothing playing.")
			return nil
		}

		if player.Paused() {
			if err := player.Pause(false); err != nil {
				return fmt.Errorf("failed to play: %w", err)
			}
			msg.SetEmbeds(embeds.Play(player.PlayingTrack().Info()))
		} else {
			if err := player.Pause(true); err != nil {
				return fmt.Errorf("failed to pause: %w", err)
			}
			msg.SetEmbeds(embeds.Pause(player.PlayingTrack().Info(), player.PlayingTrack().Info().Length-player.Position()))
		}

		msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "⏯️", "toggle", ""))
		return nil
	},
}
