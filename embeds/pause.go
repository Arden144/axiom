package embeds

import (
	"fmt"

	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/v3/lavalink"
)

func Pause(track lavalink.TrackInfo, remaining lavalink.Duration) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Blue)
	embed.SetAuthorName("Paused")
	embed.SetAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/pause_circle_outline/baseline-4x.png")
	embed.SetTitle(track.Title)
	embed.SetURL(*track.URI)
	embed.AddField("Channel", track.Author, true)
	embed.AddField("Remaining", remaining.String(), true)
	embed.SetThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/mqdefault.jpg", track.Identifier))
	return embed.Build()
}
