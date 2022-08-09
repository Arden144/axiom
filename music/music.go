package music

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgolink/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"go.uber.org/zap"
)

var (
	ctx    = context.Background()
	link   disgolink.Link
	queues map[snowflake.ID]Queue = make(map[snowflake.ID]Queue)
)

func init() {
	link = disgolink.New(bot.Client)

	_, err := link.AddNode(ctx, config.Lavalink)
	if err != nil {
		log.L.Fatal("failed to connect to lavalink", zap.Error(err))
	}
}

func GetPlayer(guildID snowflake.ID) Player {
	queue, ok := queues[guildID]
	if !ok {
		queue = newQueue()
		queues[guildID] = queue
	}

	player := Player{Queue: queue}

	player.Player = link.ExistingPlayer(guildID)
	if player.Player == nil {
		player.Player = link.Player(guildID)
		player.AddListener(newListener(player))
	}

	return player
}
