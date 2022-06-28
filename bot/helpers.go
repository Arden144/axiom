package bot

import (
	"fmt"

	"github.com/disgoorg/disgo/discord"
)

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
