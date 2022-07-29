package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/disgoorg/disgo/discord"
)

var Refresh = bot.Command{
	Create: discord.SlashCommandCreate{
		CommandName: "refresh",
		Description: "refresh",
	},
	Handler: func(_ context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		tmp := e.Bot.Config.DevGuildID
		e.Client().Caches().Guilds().ForEach(func(g discord.Guild) {
			// TODO: this is a hack
			e.Bot.Config.DevGuildID = g.ID
			for _, c := range e.Bot.Commands {
				e.Bot.AddCommands(c)
			}
		})
		e.Bot.Config.DevGuildID = tmp

		msg.SetContent("done")
		return nil
	},
}
