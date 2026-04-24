package embeds

import (
	"github.com/arden144/axiom/color"
	"github.com/disgoorg/disgo/discord"
)

func Error() discord.Embed {
	return discord.NewEmbed().
		WithColor(color.Red).
		WithAuthorName("Failed").
		WithAuthorIcon("https://github.com/material-icons/material-icons-png/raw/master/png/white/error_outline/baseline-4x.png").
		WithTitle("Unfortunately, an unexpected error has occured while trying to process your command.")
}
