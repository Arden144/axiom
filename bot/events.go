package bot

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/arden144/axiom/embeds"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func (b *Bot) OnReady(_ *events.Ready) {
	log.Print("ready")
}

var ErrCommandFailed = errors.New("expected")

func (b *Bot) OnApplicationCommandInteraction(e *events.ApplicationCommandInteractionCreate) {
	name := e.Data.CommandName()
	if c, ok := b.Commands[name]; ok {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		e := CommandEvent{e}

		defer func() {
			if r := recover(); r != nil {
				log.Printf("WARN: caught panic in command handler for %s: %v", name, r)
				if err := e.UpdateMessage(discord.NewMessageUpdateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
					log.Print("WARN: failed to send failiure acknowledgement: ", err)
				}
			}
		}()

		if err := e.DeferCreateMessage(false); err != nil {
			log.Print("WARN: failed to send command acknowledgement: ", err)
			return
		}

		msg, err := c.Handler(ctx, b, e)
		if err != nil {
			log.Print("WARN: ", err)
			if err := e.UpdateMessage(discord.NewMessageUpdateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
				log.Print("WARN: failed to send failiure acknowledgement: ", err)
			}
			return
		}

		if err := e.UpdateMessage(*msg); err != nil {
			log.Printf("WARN: failed to send response for %s: %s", name, err)
		}
	} else {
		log.Printf("WARN: %s is not a valid command name", name)
	}
}
