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
	Ctx      = context.Background()
	Commands = make(map[string]Command)
	Buttons  = make(map[string]Button)
)

func init() {
	var err error
	Client, err = disgo.New(config.Token,
		bot.WithLogger(log.W),
		bot.WithGatewayConfigOpts(
			gateway.WithIntents(gateway.IntentGuildVoiceStates),
			gateway.WithPresenceOpts(gateway.WithListeningActivity("bangers")),
		),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagVoiceStates)),
		bot.WithEventListenerFunc(OnReady),
		bot.WithEventListenerFunc(OnComponentInteraction),
		bot.WithEventListenerFunc(OnApplicationCommandInteraction),
	)
	if err != nil {
		log.L.Fatal("failed to configure bot: ", zap.Error(err))
	}

	ctx, cancel := context.WithTimeout(Ctx, 10*time.Second)
	defer cancel()

	if err := Client.OpenGateway(ctx); err != nil {
		log.L.Fatal("failed to start bot: ", zap.Error(err))
	}
}
