package bot

import (
	"context"
	"log"

	"github.com/arden144/axiom/config"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type SlashCommand = discord.SlashCommandCreate

type Command struct {
	Create  SlashCommand
	Handler func(context.Context, CommandEvent, *discord.MessageUpdateBuilder) error
}

type CommandEvent struct {
	*events.ApplicationCommandInteractionCreate
}

func ClearCommands() {
	commands, err := Client.Rest().GetGuildCommands(Client.ApplicationID(), config.Config.DevGuildID, false)
	if err != nil {
		log.Fatal("failed to get guild commands: ", err)
	}

	for _, c := range commands {
		if err := Client.Rest().DeleteGuildCommand(Client.ApplicationID(), config.Config.DevGuildID, c.ID()); err != nil {
			log.Fatal("failed to delete command: ", err)
		}
	}

	commands, err = Client.Rest().GetGlobalCommands(Client.ApplicationID(), false)
	if err != nil {
		log.Fatal("failed to get global commands: ", err)
	}

	for _, c := range commands {
		if err := Client.Rest().DeleteGlobalCommand(Client.ApplicationID(), c.ID()); err != nil {
			log.Fatal("failed to delete command: ", err)
		}
	}
}

func AddCommands(cs ...Command) {
	for _, c := range cs {
		Commands[c.Create.Name()] = c

		if _, err := Client.Rest().CreateGuildCommand(Client.ApplicationID(), config.Config.DevGuildID, c.Create); err != nil {
			log.Fatal("failed to add command: ", err)
		}
	}
}
