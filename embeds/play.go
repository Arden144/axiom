package embeds

import (
	"fmt"

	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func Play(track lavalink.TrackInfo) discord.Embed {
	return discord.NewEmbed().
		WithColor(color.Blue).
		WithAuthorName("Now Playing").
		WithAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/play_circle_outline/baseline-4x.png").
		WithTitle(track.Title).
		WithURL(*track.URI).
		AddField("Channel", track.Author, true).
		AddField("Duration", track.Length.String(), true).
		WithThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/mqdefault.jpg", track.Identifier))
}
