package embeds

import (
	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
)

func Error() discord.Embed {
	embed := discord.NewEmbedBuilder()
	embed.SetColor(color.Red)
	embed.SetTitle("Failed")
	embed.SetDescription("Unfortunately, an unexpected error has occured while trying to process your command.")
	return embed.Build()
}
