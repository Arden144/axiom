package music

import (
	"context"

	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgolink/disgolink"
	"github.com/disgoorg/disgolink/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type Music struct {
	link   disgolink.Link
	queues map[snowflake.ID]Queue
}

func New(client bot.Client, opts ...lavalink.ConfigOpt) Music {
	return Music{
		link:   disgolink.New(client, opts...),
		queues: make(map[snowflake.ID]Queue),
	}
}

func (m *Music) Connect(ctx context.Context, config lavalink.NodeConfig) error {
	_, err := m.link.AddNode(ctx, config)
	return err
}

func (m *Music) Player(guildID snowflake.ID) Player {
	queue, ok := m.queues[guildID]
	if !ok {
		queue = newQueue()
		m.queues[guildID] = queue
	}

	player := Player{Queue: &queue}

	player.Player = m.link.ExistingPlayer(guildID)
	if player.Player == nil {
		player.Player = m.link.Player(guildID)
		player.AddListener(newListener(player))
	}

	return player
}
