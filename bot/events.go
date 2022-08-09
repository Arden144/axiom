package bot

import (
	"context"
	"time"

	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgo/gateway"
	"go.uber.org/zap"
)

func OnReady(ev *events.Ready) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	Client.SetPresence(ctx, gateway.NewPresence(discord.ActivityTypeListening, "bangers", "", discord.OnlineStatusOnline, false))
	log.L.Info("ready", zap.String("username", ev.User.Username), zap.String("discriminator", ev.User.Discriminator))
}

func OnComponentInteraction(re *events.ComponentInteractionCreate) {
	id := re.ButtonInteractionData().CustomID()
	query, params, err := parseButtonID(id)
	if err != nil {
		log.L.Warn("failed to parse custom id", zap.Error(err))
		return
	}

	bt, ok := Buttons[query]
	if !ok {
		log.L.Warn("not a valid button id", zap.String("id", query))
		return
	}

	ev := ButtonEvent{re, ButtonData{params}}

	msg := discord.NewMessageCreateBuilder()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := bt.Handler(ctx, ev, msg); err != nil {
		log.L.Warn("button handler failed", zap.String("id", query), zap.Error(err))
		if err := ev.CreateMessage(discord.NewMessageCreateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.L.Warn("failed to send failiure acknowledgement: ", zap.Error(err))
		}
	}

	if err := ev.CreateMessage(msg.Build()); err != nil {
		log.L.Warn("failed to send response", zap.String("id", query), zap.Error(err))
	}
}

func OnApplicationCommandInteraction(re *events.ApplicationCommandInteractionCreate) {
	if err := re.DeferCreateMessage(false); err != nil {
		log.L.Warn("failed to send command acknowledgement: ", zap.Error(err))
		return
	}

	ev := CommandEvent{re}

	name := ev.Data.CommandName()
	c, ok := Commands[name]
	if !ok {
		log.L.Warn("not a valid command name", zap.String("command", name))
		return
	}

	msg := discord.NewMessageUpdateBuilder()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := c.Handler(ctx, ev, msg); err != nil {
		log.L.Warn("button handler failed", zap.String("command", name), zap.Error(err))
		if err := ev.UpdateMessage(discord.NewMessageUpdateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.L.Warn("failed to send failiure acknowledgement: ", zap.Error(err))
		}
		return
	}

	if err := ev.UpdateMessage(msg.Build()); err != nil {
		log.L.Warn("failed to send response", zap.String("command", name), zap.Error(err))
	}
}
