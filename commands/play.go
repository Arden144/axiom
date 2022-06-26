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
		CommandName: "play",
		Description: "play",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "song",
				Description: "song",
				Required:    true,
			},
		},
	},
	Handler: func(ctx context.Context, b *bot.Bot, e bot.CommandEvent) (*discord.MessageUpdate, error) {
		song := e.SlashCommandInteractionData().String("song")
		player := b.Music.Player(*e.GuildID())

		voice, ok := b.Client.Caches().VoiceStates().Get(*e.GuildID(), e.User().ID)
		if !ok {
			return e.Reply("Not in a voice channel")
		}

		if player.ChannelID() != voice.ChannelID {
			if err := b.Client.Connect(ctx, *e.GuildID(), *voice.ChannelID); err != nil {
				return e.Fatal("failed to join channel", err)
			}
		}

		tracks, err := player.Search(ctx, song)
		if err == music.ErrNotFound {
			return e.Reply("not found")
		} else if err != nil {
			return e.Fatal("failed to search", err)
		}

		track := tracks[0]

		if player.Playing() {
			player.Enqueue(track)
			return e.Reply(fmt.Sprint("Queued ", track.Info().Title))
		} else {
			if err := player.Play(track); err != nil {
				return e.Fatal("failed to play", err)
			}
			return e.ReplyEmbed(embeds.Play(track.Info()))
		}
	},
}
