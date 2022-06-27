package embeds

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/lavalink"
)

func Play(track lavalink.AudioTrackInfo) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetAuthorName("Now playing")
	embed.SetTitle(track.Title)
	embed.SetURL(*track.URI)
	embed.SetField(0, "Channel", track.Author, true)
	embed.SetField(1, "Duration", track.Length.String(), true)
	embed.SetThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/default.jpg", track.Identifier))
	return embed.Build()
}
