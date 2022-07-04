package embeds

import (
	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
)

func Error() discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Red)
	embed.SetAuthorName("Failed")
	embed.SetAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/error_outline/baseline.png")
	embed.SetTitle("Unfortunately, an unexpected error has occured while trying to process your command.")
	return embed.Build()
}
