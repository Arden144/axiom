package bot

import (
	"context"
	"time"

	"github.com/arden144/axiom/embeds"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/disgolink/v3/lavalink"
	"go.uber.org/zap"
)

func OnReady(ev *events.Ready) {
	log.L.Info("bot ready", zap.String("username", ev.User.Username), zap.String("discriminator", ev.User.Discriminator))
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
	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()

	if err := bt.Handler(ctx, ev, msg); err != nil {
		log.L.Warn("button handler failed", zap.String("id", query), zap.Error(err))
		if err := ev.CreateMessage(discord.NewMessageCreateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.L.Warn("failed to send failiure acknowledgement", zap.Error(err))
		}
	}

	if err := ev.CreateMessage(msg.Build()); err != nil {
		log.L.Warn("failed to send response", zap.String("id", query), zap.Error(err))
	}
}

func OnApplicationCommandInteraction(re *events.ApplicationCommandInteractionCreate) {
	if err := re.DeferCreateMessage(false); err != nil {
		log.L.Warn("failed to send command acknowledgement", zap.Error(err))
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
	ctx, cancel := context.WithTimeout(Ctx, 5*time.Second)
	defer cancel()

	if err := c.Handler(ctx, ev, msg); err != nil {
		log.L.Warn("command handler failed", zap.String("command", name), zap.Error(err))
		if err := ev.UpdateMessage(discord.NewMessageUpdateBuilder().SetEmbeds(embeds.Error()).Build()); err != nil {
			log.L.Warn("failed to send failiure acknowledgement", zap.Error(err))
		}
		return
	}

	if err := ev.UpdateMessage(msg.Build()); err != nil {
		log.L.Warn("failed to send response", zap.String("command", name), zap.Error(err))
	}
}

func OnVoiceStateUpdate(event *events.GuildVoiceStateUpdate) {
	if event.VoiceState.UserID != Client.ApplicationID() {
		return
	}
	Link.OnVoiceStateUpdate(context.TODO(), event.VoiceState.GuildID, event.VoiceState.ChannelID, event.VoiceState.SessionID)
}

func OnVoiceServerUpdate(event *events.VoiceServerUpdate) {
	Link.OnVoiceServerUpdate(context.TODO(), event.GuildID, event.Token, *event.Endpoint)
}

func OnTrackEnd(player disgolink.Player, event lavalink.TrackEndEvent) {
	if !event.Reason.MayStartNext() {
		return
	}

	p := GetPlayer(player.GuildID())

	if err := p.Next(); err != nil {
		log.L.Warn("failed to play", zap.Error(err))
	}
}
