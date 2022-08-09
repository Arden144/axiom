package bot

import (
	"context"

	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"go.uber.org/zap"
)

type SlashCommand = discord.SlashCommandCreate

type CommandEvent struct {
	*events.ApplicationCommandInteractionCreate
}

type Command struct {
	Create  SlashCommand
	Handler func(context.Context, CommandEvent, *discord.MessageUpdateBuilder) error
}

func ClearCommands() {
	commands, err := Client.Rest().GetGuildCommands(Client.ApplicationID(), config.DevGuildID, false)
	if err != nil {
		log.L.Fatal("failed to get guild commands", zap.Error(err))
	}

	for _, c := range commands {
		if err := Client.Rest().DeleteGuildCommand(Client.ApplicationID(), config.DevGuildID, c.ID()); err != nil {
			log.L.Fatal("failed to delete command", zap.Error(err))
		}
	}

	commands, err = Client.Rest().GetGlobalCommands(Client.ApplicationID(), false)
	if err != nil {
		log.L.Fatal("failed to get global commands", zap.Error(err))
	}

	for _, c := range commands {
		if err := Client.Rest().DeleteGlobalCommand(Client.ApplicationID(), c.ID()); err != nil {
			log.L.Fatal("failed to delete command", zap.Error(err))
		}
	}
}

func AddCommands(cs ...Command) {
	for _, c := range cs {
		Commands[c.Create.Name] = c

		if _, err := Client.Rest().CreateGuildCommand(Client.ApplicationID(), config.DevGuildID, c.Create); err != nil {
			log.L.Fatal("failed to add command: ", zap.Error(err))
		}
	}
}
