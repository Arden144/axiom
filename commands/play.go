package commands

import (
	"context"
	"fmt"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

func voiceChannelId(b *bot.Bot, guildID, userID snowflake.ID) (*snowflake.ID, bool) {
	state, ok := b.Client.Caches().VoiceStates().GroupFindFirst(guildID, func(_ snowflake.ID, state discord.VoiceState) bool {
		return state.UserID == userID
	})
	return state.ChannelID, ok
}

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
	Handler: func(ctx context.Context, b *bot.Bot, e *bot.CommandEvent) bot.Message {
		song := e.SlashCommandInteractionData().String("song")

		channelID, ok := voiceChannelId(b, *e.GuildID(), e.Member().User.ID)
		if !ok {
			return e.Reply("Not in a voice channel")
		}

		if err := b.Client.Connect(ctx, *e.GuildID(), *channelID); err != nil {
			fmt.Print("WARM: failed to join channel: ", err)
			return e.Reply("failed")
		}

		player := b.Music.Player(*e.GuildID())
		info, err := player.Play(ctx, song)
		if err == music.ErrNotFound {
			return e.Reply("not found")
		} else if err != nil {
			fmt.Print("WARN: failed to play song: ", err)
			return e.Reply("failed")
		}

		return e.Reply(fmt.Sprint("Now playing ", info.Title))
	},
}
