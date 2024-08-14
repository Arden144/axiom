package commands

import (
	"context"
	"fmt"
	"strings"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/music"
	"github.com/arden144/axiom/search"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

var Play = bot.Command{
	Create: bot.SlashCommand{
		Name:        "play",
		Description: "Play a song in your current voice channel",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "song",
				Description: "The name of the song you'd like to play",
				Required:    true,
			},
		},
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		song := e.SlashCommandInteractionData().String("song")
		player := bot.GetPlayer(*e.GuildID())

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

		info, err := search.Search(ctx, song)
		if err != nil {
			return fmt.Errorf("search failed: %w", err)
		}

		tracks, err := player.Search(ctx, fmt.Sprintf("%v - %v", info.Artists[0].Name, info.Name))
		if err == music.ErrNotFound {
			msg.SetContent("not found")
			return nil
		} else if err != nil {
			return fmt.Errorf("failed to search: %w", err)
		}

		track := tracks[0]
		for _, v := range tracks {
			title := strings.ToLower(v.Info.Title)
			if !strings.Contains(title, "official video") && !strings.Contains(title, "music video") {
				track = v
				break
			}
		}

		if player.Playing() {
			length := player.Track().Info.Length - player.Position() + player.Remaining()
			player.Enqueue(track)
			msg.SetEmbeds(embeds.Queue(track.Info, length))
		} else {
			if err := player.Update(ctx, lavalink.WithTrack(track)); err != nil {
				return fmt.Errorf("failed to play: %w", err)
			}
			msg.SetEmbeds(embeds.Play(track.Info))
		}

		msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "⏯️", "toggle", "", 0))
		return nil
	},
}
