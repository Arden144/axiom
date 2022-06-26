package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
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

		if player.Playing() {
			player.Enqueue(tracks[0])
			return e.Reply(fmt.Sprint("Queued ", tracks[0].Info().Title))
		} else {
			e.Fatal("failed to play", player.Play(tracks[0]))
			return e.Reply(fmt.Sprint("Now playing ", tracks[0].Info().Title))
		}
	},
}
