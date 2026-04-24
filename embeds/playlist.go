package embeds

import (
	"strconv"

	"github.com/arden144/axiom/color"
	"github.com/arden144/axiom/search"
	"github.com/disgoorg/disgo/discord"
)

func Playlist(playlist search.Playlist, queued int) discord.Embed {
	embed := discord.NewEmbed().
		WithColor(color.Blue).
		WithAuthorName("Playlist added to queue").
		WithAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/control_point/baseline-4x.png").
		WithTitle(playlist.Name).
		WithURL(playlist.ExternalURLs.Spotify).
		AddField("Total tracks", strconv.Itoa(len(playlist.Tracks.Items)), true).
		AddField("Queued tracks", strconv.Itoa(queued), true)
	if (len(playlist.Images)) > 0 {
		embed = embed.WithThumbnail(playlist.Images[0].Url)
	}
	return embed
}
