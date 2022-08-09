package bot

import (
	"context"
	"time"

	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/gateway"
	"go.uber.org/zap"
)

var (
	Client   bot.Client
	Commands map[string]Command = make(map[string]Command)
	Buttons  map[string]Button  = make(map[string]Button)
)

var ctx = context.Background()

func init() {
	var err error
	Client, err = disgo.New(config.Token,
		bot.WithLogger(log.W),
		bot.WithGatewayConfigOpts(gateway.WithIntents(gateway.IntentGuildVoiceStates)),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagVoiceStates)),
		bot.WithEventListenerFunc(OnReady),
		bot.WithEventListenerFunc(OnComponentInteraction),
		bot.WithEventListenerFunc(OnApplicationCommandInteraction),
	)
	if err != nil {
		log.L.Fatal("failed to configure bot: ", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := Client.OpenGateway(ctx); err != nil {
		log.L.Fatal("failed to start bot: ", zap.Error(err))
	}
}
