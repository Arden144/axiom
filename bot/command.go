package bot

import (
	"context"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type SlashCommand = discord.SlashCommandCreate

type CommandEvent struct {
	*events.ApplicationCommandInteractionCreate
}

type Command struct {
	Create  SlashCommand
	Handler func(context.Context, CommandEvent, *discord.MessageUpdateBuilder) error
}
