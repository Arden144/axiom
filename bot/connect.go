package bot

import (
	"context"
	"log"
	"time"
)

const timeout = 10 * time.Second

func (b *Bot) connectDiscord(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := b.Client.ConnectGateway(ctx); err != nil {
		log.Fatal("failed to start bot: ", err)
	}
}

func (b *Bot) connectLavalink(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := b.Music.Connect(ctx, b.Config.Lavalink); err != nil {
		log.Fatal("failed to add lavalink node: ", err)
	}
}
