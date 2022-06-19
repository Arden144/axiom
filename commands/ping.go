package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
)

var Ping = bot.Command{
	Create: bot.SlashCommand{
		CommandName: "ping",
		Description: "Ping",
	},
	Handler: func(_ context.Context, b *bot.Bot, e *bot.CommandEvent) bot.Message {
		return e.Reply("pong")
	},
}
