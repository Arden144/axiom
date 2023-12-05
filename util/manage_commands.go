package util

import (
	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/snowflake/v2"
	"go.uber.org/zap"
)

func ClearCommands(from snowflake.ID) {
	commands, err := bot.Client.Rest().GetGuildCommands(bot.Client.ApplicationID(), from, false)
	if err != nil {
		log.L.Fatal("failed to get guild commands", zap.Error(err))
	}

	for _, c := range commands {
		if err := bot.Client.Rest().DeleteGuildCommand(bot.Client.ApplicationID(), from, c.ID()); err != nil {
			log.L.Fatal("failed to delete command", zap.Error(err))
		}
	}

	// commands, err = Client.Rest().GetGlobalCommands(Client.ApplicationID(), false)
	// if err != nil {
	// 	log.L.Fatal("failed to get global commands", zap.Error(err))
	// }

	// for _, c := range commands {
	// 	if err := Client.Rest().DeleteGlobalCommand(Client.ApplicationID(), c.ID()); err != nil {
	// 		log.L.Fatal("failed to delete command", zap.Error(err))
	// 	}
	// }
}

func AddCommands(to snowflake.ID, cs ...bot.Command) {
	for _, c := range cs {
		bot.Commands[c.Create.Name] = c

		if _, err := bot.Client.Rest().CreateGuildCommand(bot.Client.ApplicationID(), to, c.Create); err != nil {
			log.L.Fatal("failed to add command", zap.Error(err))
		}
	}
}
