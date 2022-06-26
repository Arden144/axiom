package bot

import (
	"context"
	"fmt"
	"log"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

type SlashCommand = discord.SlashCommandCreate

type Command struct {
	Create  SlashCommand
	Handler func(context.Context, *Bot, CommandEvent) (*discord.MessageUpdate, error)
}

type CommandEvent struct {
	*events.ApplicationCommandInteractionCreate
}

func (e *CommandEvent) Reply(content string) (*discord.MessageUpdate, error) {
	embed := discord.NewEmbedBuilder().SetDescription(content).Build()
	msg := discord.NewMessageUpdateBuilder().SetEmbeds(embed).Build()
	return &msg, nil
}

func (e *CommandEvent) ReplyEmbed(embed discord.Embed) (*discord.MessageUpdate, error) {
	msg := discord.NewMessageUpdateBuilder().SetEmbeds(embed).Build()
	return &msg, nil
}

func (e *CommandEvent) Fatal(msg string, err error) (*discord.MessageUpdate, error) {
	return nil, fmt.Errorf("%s: %w", msg, err)

}

func (e *CommandEvent) UpdateMessage(messageUpdate discord.MessageUpdate) error {
	_, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate)
	return err
}

func (b *Bot) ClearCommands() {
	commands, err := b.Client.Rest().GetGuildCommands(b.Client.ApplicationID(), b.Config.DevGuildID, false)
	if err != nil {
		log.Fatal("failed to get guild commands: ", err)
	}

	for _, c := range commands {
		if err := b.Client.Rest().DeleteGuildCommand(b.Client.ApplicationID(), b.Config.DevGuildID, c.ID()); err != nil {
			log.Fatal("failed to delete command: ", err)
		}
	}

	commands, err = b.Client.Rest().GetGlobalCommands(b.Client.ApplicationID(), false)
	if err != nil {
		log.Fatal("failed to get global commands: ", err)
	}

	for _, c := range commands {
		if err := b.Client.Rest().DeleteGlobalCommand(b.Client.ApplicationID(), c.ID()); err != nil {
			log.Fatal("failed to delete command: ", err)
		}
	}
}

func (b *Bot) AddCommand(c Command) {
	b.Commands[c.Create.Name()] = c

	if _, err := b.Client.Rest().CreateGuildCommand(b.Client.ApplicationID(), b.Config.DevGuildID, c.Create); err != nil {
		log.Fatal("failed to add command: ", err)
	}
}

func (b *Bot) AddCommands(cs ...Command) {
	for _, c := range cs {
		b.AddCommand(c)
	}
}
