package embeds

import (
	"fmt"

	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/lavalink"
)

func Play(track lavalink.AudioTrackInfo) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Blue)
	embed.SetAuthorName("Now playing")
	embed.SetTitle(track.Title)
	embed.SetURL(*track.URI)
	embed.AddField("Channel", track.Author, true)
	embed.AddField("Duration", track.Length.String(), true)
	embed.SetThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/default.jpg", track.Identifier))
	return embed.Build()
}
