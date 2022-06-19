package music

import (
	"context"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgolink/disgolink"
	"github.com/disgoorg/disgolink/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type Music struct {
	link disgolink.Link
}

func New(client bot.Client, opts ...lavalink.ConfigOpt) Music {
	return Music{disgolink.New(client, opts...)}
}

func (m *Music) Connect(ctx context.Context, config lavalink.NodeConfig) error {
	_, err := m.link.AddNode(ctx, config)
	return err
}

func (m *Music) Player(guildID snowflake.ID) *Player {
	return &Player{
		link:   m.link,
		player: m.link.Player(guildID),
		queue:  make(Queue, 0),
	}
}
