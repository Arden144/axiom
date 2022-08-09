package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
)

var Play = bot.Command{
	Create: bot.SlashCommand{
		Name:        "play",
		Description: "play",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "song",
				Description: "song",
				Required:    true,
			},
		},
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		song := e.SlashCommandInteractionData().String("song")
		player := music.GetPlayer(*e.GuildID())

		voice, ok := bot.Client.Caches().VoiceStates().Get(*e.GuildID(), e.User().ID)
		if !ok {
			msg.SetContent("Not in a voice channel")
			return nil
		}

		if player.ChannelID() != voice.ChannelID {
			if err := bot.Client.Connect(ctx, *e.GuildID(), *voice.ChannelID); err != nil {
				return fmt.Errorf("failed to join channel: %w", err)
			}
		}

		tracks, err := player.Search(ctx, song)
		if err == music.ErrNotFound {
			msg.SetContent("not found")
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to search: %w", err)
		}

		track := tracks[0]

		if player.Playing() {
			player.Enqueue(track)
			msg.SetContent(fmt.Sprint("Queued ", track.Info().Title))
		} else {
			if err := player.Play(track); err != nil {
				return fmt.Errorf("failed to play: %w", err)
			}
			msg.SetEmbeds(embeds.Play(track.Info()))
		}

		msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "Pause", "pause", ""))
		return nil
	},
}
