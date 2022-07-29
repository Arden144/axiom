package bot

import (
	"context"
	"log"

	"github.com/arden144/axiom/config"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

var (
	Client   bot.Client
	Commands map[string]Command = make(map[string]Command)
	Buttons  map[string]Button  = make(map[string]Button)
)

var ctx = context.Background()

func init() {
	var err error
	Client, err = disgo.New(config.Config.Token,
		bot.WithGatewayConfigOpts(gateway.WithGatewayIntents(discord.GatewayIntentGuildVoiceStates)),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagVoiceStates)),
		bot.WithEventListenerFunc(OnReady),
		bot.WithEventListenerFunc(OnComponentInteraction),
		bot.WithEventListenerFunc(OnApplicationCommandInteraction),
	)
	if err != nil {
		log.Fatal("failed to configure bot: ", err)
	}

	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	if err := Client.ConnectGateway(ctx); err != nil {
		log.Fatal("failed to start bot: ", err)
	}
}

// func (b *Bot) Start() {
// 	if b.Client == nil {
// 		log.Fatal("call bot.Setup() before bot.Start()")
// 	}

// 	ctx := context.Background()

// 	log.Print("starting")
// 	b.connectDiscord(ctx)
// 	b.connectLavaLink(ctx)

// 	utility.OnSignal(os.Interrupt)

// 	log.Print("stopping")
// 	b.disconnectDiscord(ctx)
// 	b.disconnectLavaLink()
// 	log.Print("done")
// }
