package buttons

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var Toggle = bot.Button{
	Query: "toggle",
	Handler: func(ctx context.Context, e bot.ButtonEvent, msg *discord.MessageCreateBuilder) error {
		player := bot.GetPlayer(*e.GuildID())

		if !player.Playing() {
			msg.SetContent("There's nothing playing.")
			return nil
		}

		if player.Paused() {
			if err := player.Update(ctx, lavalink.WithPaused(false)); err != nil {
				return fmt.Errorf("failed to play: %w", err)
			}
			msg.SetEmbeds(embeds.Play(player.Track().Info))
		} else {
			if err := player.Update(ctx, lavalink.WithPaused(true)); err != nil {
				return fmt.Errorf("failed to pause: %w", err)
			}
			msg.SetEmbeds(embeds.Pause(player.Track().Info, player.Track().Info.Length-player.Position()))
		}

		msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "⏯️", "toggle", "", 0))
		return nil
	},
}
