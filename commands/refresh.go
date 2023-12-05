package commands

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/util"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/snowflake/v2"
)

var Refresh = bot.Command{
	Create: bot.SlashCommand{
		Name:        "refresh",
		Description: "Clears and readds all commands to a server.",
		Options: []discord.ApplicationCommandOption{
			discord.ApplicationCommandOptionString{
				Name:        "GuildID",
				Description: "The ID of the server to refresh.",
				Required:    true,
			},
		},
	},
	Handler: func(ctx context.Context, e bot.CommandEvent, msg *discord.MessageUpdateBuilder) error {
		guildIDString := e.SlashCommandInteractionData().String("GuildID")
		guildID, err := snowflake.Parse(guildIDString)
		if err != nil {
			msg.SetContent("Not a valid Guild ID.")
			return nil
		}

		util.ClearCommands(guildID)
		util.AddCommands(guildID, Play, PlayLink, Skip, Disconnect, Pause, Resume)

		msg.SetContent("Done.")
		return nil
	},
}
