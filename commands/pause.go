package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var Pause = bot.Command{
	Create: bot.SlashCommand{
		Name:        "pause",
		Description: "pause",
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdate) error {
		player := bot.GetPlayer(*e.GuildID())

		if !player.Playing() {
			*msg = msg.WithContent("nothing to pause")
			return nil
		}

		if player.Paused() {
			*msg = msg.WithContent("already paused")
			return nil
		}

		if err := player.Update(ctx, lavalink.WithPaused(true)); err != nil {
			return fmt.Errorf("failed to pause: %w", err)
		}

		*msg = msg.WithEmbeds(embeds.Pause(player.Track().Info, player.Track().Info.Length-player.Position()))
		return nil
	},
}
