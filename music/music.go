package music

import (
	"context"

	"github.com/arden144/axiom/bot"
	"github.com/arden144/axiom/config"
	"github.com/disgoorg/disgolink/disgolink"
	"github.com/disgoorg/snowflake/v2"
)

var link disgolink.Link
var queues map[snowflake.ID]Queue = make(map[snowflake.ID]Queue)

var ctx = context.Background()

func init() {
	link = disgolink.New(bot.Client)
	link.AddNode(ctx, config.Config.Lavalink)
}

// func (m *Music) Disconnect() {
// 	for _, n := range m.link.Nodes() {
// 		m.link.RemoveNode(n.Name())
// 	}
// }

func Player(guildID snowflake.ID) PlayerType {
	queue, ok := queues[guildID]
	if !ok {
		queue = newQueue()
		queues[guildID] = queue
	}

	player := PlayerType{Queue: queue}

	player.Player = link.ExistingPlayer(guildID)
	if player.Player == nil {
		player.Player = link.Player(guildID)
		player.AddListener(newListener(player))
	}

	return player
}
