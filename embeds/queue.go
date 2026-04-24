package embeds

import (
	"fmt"

	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func Queue(track lavalink.TrackInfo, remaining lavalink.Duration) discord.Embed {
	return discord.NewEmbed().
		WithColor(color.Yellow).
		WithAuthorName("Added to Queue").
		WithAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/control_point/baseline-4x.png").
		WithTitle(track.Title).
		WithURL(*track.URI).
		AddField("Channel", track.Author, true).
		AddField("Playing in", remaining.String(), true).
		WithThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/mqdefault.jpg", track.Identifier))
}
