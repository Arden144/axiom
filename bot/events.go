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

func (b *Bot) OnComponentInteraction(re *events.ComponentInteractionCreate) {
	id := re.ButtonInteractionData().CustomID()
	query, params, err := parse(id)
	if err != nil {
		log.Print("WARN: failed to parse custom id: ", err)
		return
	}

	bt, ok := b.Buttons[query]
	if !ok {
		log.Printf("WARN: %s is not a valid button id", query)
		return
	}

	e := ButtonEvent{re, ButtonData{params}, b}

	msg := discord.NewMessageCreateBuilder()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := bt.Handler(ctx, e, msg); err != nil {
		log.Printf("WARN: button handler for %v failed: %v", query, err)
		if err := e.CreateMessage(discord.NewMessageCreateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.Print("WARN: failed to send failiure acknowledgement: ", err)
		}
	}

	if err := e.CreateMessage(msg.Build()); err != nil {
		log.Printf("WARN: failed to send response for %v: %v", query, err)
	}
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
