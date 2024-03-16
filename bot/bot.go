package bot

import (
	"context"
	"time"

	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"github.com/arden144/axiom/music"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"go.uber.org/zap"
)

var (
	Client   bot.Client
	Link     disgolink.Client
	Ctx      = context.Background()
	Commands = make(map[string]Command)
	Buttons  = make(map[string]Button)
	queues   = make(map[snowflake.ID]music.Queue)
)

func init() {
	var err error
	Client, err = disgo.New(config.Token,
		bot.WithLogger(log.SL),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(gateway.IntentGuildVoiceStates),
			gateway.WithPresenceOpts(gateway.WithListeningActivity("bangers")),
		),
		bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagVoiceStates)),
		bot.WithEventListenerFunc(OnReady),
		bot.WithEventListenerFunc(OnComponentInteraction),
		bot.WithEventListenerFunc(OnApplicationCommandInteraction),
		bot.WithEventListenerFunc(OnVoiceServerUpdate),
		bot.WithEventListenerFunc(OnVoiceStateUpdate),
	)
	if err != nil {
		log.L.Fatal("failed to configure bot: ", zap.Error(err))
	}

	Link = disgolink.New(Client.ID(),
		disgolink.WithListenerFunc(OnTrackEnd),
	)

	_, err = Link.AddNode(Ctx, config.Lavalink)
	if err != nil {
		log.L.Fatal("failed to connect to lavalink", zap.Error(err))
	}

	log.L.Info("connected to lavalink", zap.String("address", config.Lavalink.Address))

	ctx, cancel := context.WithTimeout(Ctx, 10*time.Second)
	defer cancel()

	if err := Client.OpenGateway(ctx); err != nil {
		log.L.Fatal("failed to start bot: ", zap.Error(err))
	}
}

func getQueue(guildID snowflake.ID) music.Queue {
	queue, ok := queues[guildID]
	if !ok {
		queue = music.NewQueue()
		queues[guildID] = queue
	}

	return queue
}

func GetPlayer(guildID snowflake.ID) music.Player {
	player := Link.Player(guildID)
	queue := getQueue(guildID)
	return music.Player{Player: player, Queue: queue}
}
