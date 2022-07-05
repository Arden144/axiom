package bot

import (
	"context"
	"log"
	"time"

	"github.com/arden144/axiom/embeds"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

func (b *Bot) OnReady(_ *events.Ready) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	b.Client.SetPresence(ctx, discord.NewListeningPresence("bangers", discord.OnlineStatusOnline, false))
	log.Print("ready")
}

func (b *Bot) OnApplicationCommandInteraction(re *events.ApplicationCommandInteractionCreate) {
	if err := re.DeferCreateMessage(false); err != nil {
		log.Print("WARN: failed to send command acknowledgement: ", err)
		return
	}

	e := CommandEvent{re, b}

	name := e.Data.CommandName()
	c, ok := b.Commands[name]
	if !ok {
		log.Printf("WARN: %s is not a valid command name", name)
		return
	}

	msg := discord.NewMessageUpdateBuilder()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Handler(ctx, e, msg); err != nil {
		log.Printf("WARN: command handler for %v failed: %v", name, err)
		if err := e.UpdateMessage(discord.NewMessageUpdateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.Print("WARN: failed to send failiure acknowledgement: ", err)
		}
		return
	}

	if err := e.UpdateMessage(msg.Build()); err != nil {
		log.Printf("WARN: failed to send response for %v: %v", name, err)
	}
}
