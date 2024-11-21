package embeds

import (
	"strconv"

	"github.com/arden144/axiom/color"
	"github.com/arden144/axiom/search"
	"github.com/disgoorg/disgo/discord"
)

func Playlist(playlist search.Playlist, queued int) discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Blue)
	embed.SetAuthorName("Playlist added to queue")
	embed.SetAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/control_point/baseline-4x.png")
	embed.SetTitle(playlist.Name)
	embed.SetURL(playlist.ExternalURLs.Spotify)
	embed.AddField("Total tracks", strconv.Itoa(len(playlist.Tracks.Items)), true)
	embed.AddField("Queued tracks", strconv.Itoa(queued), true)
	if (len(playlist.Images)) > 0 {
		embed.SetThumbnail(playlist.Images[0].Url)
	}
	return embed.Build()
}
