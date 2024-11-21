package commands

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/log"
	"github.com/arden144/axiom/music"
	"github.com/arden144/axiom/search"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"go.uber.org/zap"
)

var Playlist = bot.Command{Create: bot.SlashCommand{
	Name:        "playlist",
	Description: "Add all songs in a playlist to the queue.",
	Options: []discord.ApplicationCommandOption{
		discord.ApplicationCommandOptionString{
			Name:        "name",
			Description: "The name of the playlist you'd like to add",
			Required:    true,
		},
	},
}, Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
	name := e.SlashCommandInteractionData().String("name")
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

	playlist, err := search.SearchPlaylist(ctx, name)
	if err != nil {
		return fmt.Errorf("search failed: %w", err)
	}

	var errs []error

	for _, track := range playlist.Tracks.Items {
		if err := FindAndPlay(ctx, player, track.Track); err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) == len(playlist.Tracks.Items) {
		return fmt.Errorf("all tracks failed to play: %w", errors.Join(errs...))
	}

	if len(errs) > 0 {
		log.L.Warn("some tracks failed to play", zap.Error(errors.Join(errs...)))
	}

	msg.SetEmbeds(embeds.Playlist(*playlist, len(playlist.Tracks.Items)-len(errs)))
	msg.AddActionRow(discord.NewButton(discord.ButtonStylePrimary, "⏯️", "toggle", "", 0))
	return nil
}}

func FindAndPlay(ctx context.Context, player music.Player, info search.Track) error {
	tracks, err := player.Search(ctx, fmt.Sprintf("%v - %v", info.Artists[0].Name, info.Name))
	if err == music.ErrNotFound {
		return fmt.Errorf("%v - %v not found", info.Artists[0].Name, info.Name)
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
		player.Enqueue(track)
	} else {
		if err := player.Update(ctx, lavalink.WithTrack(track)); err != nil {
			return fmt.Errorf("failed to play: %w", err)
		}
	}

	return nil
}
