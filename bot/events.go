package bot

import (
	"context"
	"log"
	"time"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func (b *Bot) OnReady(_ *events.Ready) {
	log.Print("ready")
}

func (b *Bot) OnApplicationCommandInteraction(e *events.ApplicationCommandInteractionCreate) {
	name := e.Data.CommandName()
	if c, ok := b.Commands[name]; ok {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		msg := c.Handler(ctx, b, &CommandEvent{e})

		if err := e.CreateMessage(discord.MessageCreate(msg)); err != nil {
			log.Printf("WARN: failed to send response for %s: %s", name, err)
		}
	} else {
		log.Printf("WARN: %s is not a valid command name", name)
	}
}
