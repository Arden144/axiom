package bot

import (
	"context"
	"log"
	"os"

	"github.com/arden144/axiom/config"
	"github.com/arden144/axiom/music"
	"github.com/arden144/axiom/utility"
	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
)

type Bot struct {
	Client   bot.Client
	Config   config.Config
	Music    music.Music
	Commands map[string]Command
	Buttons  map[string]Button
}

func New(config config.Config) Bot {
	return Bot{
		Config:   config,
		Commands: make(map[string]Command),
		Buttons:  make(map[string]Button),
	}
}

func (b *Bot) Setup() {
	var err error

	b.Client, err = disgo.New(b.Config.Token,
		bot.WithGatewayConfigOpts(gateway.WithGatewayIntents(discord.GatewayIntentGuildVoiceStates, discord.GatewayIntentGuilds)),
		bot.WithCacheConfigOpts(cache.WithCacheFlags(cache.FlagVoiceStates, cache.FlagGuilds)),
		bot.WithEventListenerFunc(b.OnReady),
		bot.WithEventListenerFunc(b.OnComponentInteraction),
		bot.WithEventListenerFunc(b.OnApplicationCommandInteraction),
	)
	if err != nil {
		log.Fatal("failed to configure bot: ", err)
	}

	b.Music = music.New(b.Client)
}

func (b *Bot) Start() {
	if b.Client == nil {
		log.Fatal("call bot.Setup() before bot.Start()")
	}

	ctx := context.Background()

	log.Print("starting")
	b.connectDiscord(ctx)
	b.connectLavaLink(ctx)

	utility.OnSignal(os.Interrupt)

	log.Print("stopping")
	b.disconnectDiscord(ctx)
	b.disconnectLavaLink()
	log.Print("done")
}
