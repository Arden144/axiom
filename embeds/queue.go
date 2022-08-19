package embeds

import (
	"fmt"

	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/lavalink"
)

func Queue(track lavalink.AudioTrackInfo, remaining lavalink.Duration) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Yellow)
	embed.SetAuthorName("Added to Queue")
	embed.SetAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/control_point/baseline-4x.png")
	embed.SetTitle(track.Title)
	embed.SetURL(*track.URI)
	embed.AddField("Channel", track.Author, true)
	embed.AddField("Playing in", remaining.String(), true)
	embed.SetThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/mqdefault.jpg", track.Identifier))
	return embed.Build()
}
