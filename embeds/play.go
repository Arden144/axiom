package embeds

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgolink/lavalink"
)

func Play(track lavalink.AudioTrackInfo) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetTitle(track.Title)
	embed.SetThumbnail(fmt.Sprintf("https://img.youtube.com/vi/%v/default.jpg", track.Identifier))
	embed.SetFields(discord.EmbedField{
		Name:  "Channel",
		Value: track.Author,
	}, discord.EmbedField{
		Name:  "Duration",
		Value: track.Length.String(),
	})
	return embed.Build()
}
