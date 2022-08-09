package bot

import "github.com/disgoorg/disgo/discord"

func (e *CommandEvent) UpdateMessage(messageUpdate discord.MessageUpdate) error {
	_, err := e.Client().Rest().UpdateInteractionResponse(e.ApplicationID(), e.Token(), messageUpdate)
	return err
}
