package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var PlayLink = bot.Command{
	Create: bot.SlashCommand{
		Name:        "playlink",
		Description: "Play a song in your current voice channel",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "url",
				Description: "A link to the YouTube video you'd like to play",
				Required:    true,
			},
		},
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		url := e.SlashCommandInteractionData().String("url")
		player := music.GetPlayer(*e.GuildID())

		voice, ok := bot.Client.Caches().VoiceState(*e.GuildID(), e.User().ID)
		if !ok {
			msg.SetContent("Not in a voice channel")
			return nil
		}

		if player.ChannelID() != voice.ChannelID {
			if err := bot.Client.UpdateVoiceState(ctx, *e.GuildID(), voice.ChannelID, false, true); err != nil {
				return fmt.Errorf("failed to join channel: %w", err)
			}
		}

		track, err := player.ResolveUrl(ctx, url)
		if err == music.ErrNotFound {
			msg.SetContent("not found")
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to resolve youtube url: %w", err)
		}

		if player.Playing() {
			length := player.PlayingTrack().Info().Length - player.Position() + player.Remaining()
			player.Enqueue(track)
			msg.SetEmbeds(embeds.Queue(track.Info(), length))
		} else {
			if err := player.Play(track); err != nil {
				return fmt.Errorf("failed to play: %w", err)
			}
			msg.SetEmbeds(embeds.Play(track.Info()))
		}

		msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "⏯️", "toggle", ""))
		return nil
	},
}
