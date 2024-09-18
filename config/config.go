package config

import (
	"encoding/json"
	"os"

	"github.com/arden144/axiom/log"
	"github.com/disgoorg/disgolink/v3/disgolink"
	"github.com/disgoorg/snowflake/v2"
	"go.uber.org/zap"
)

const path = "config/config.json"

var (
	Lavalink   disgolink.NodeConfig
	Spotify    SpotifyConfig
	Token      string
	DevGuildID snowflake.ID
)

type SpotifyConfig struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type config struct {
	Lavalink   disgolink.NodeConfig `json:"lavalink"`
	Spotify    SpotifyConfig        `json:"spotify"`
	Token      string               `json:"token"`
	DevGuildID snowflake.ID         `json:"devGuildId"`
}

func init() {
	file, err := os.Open(path)
	if err != nil {
		log.L.Fatal("failed to read config", zap.Error(err))
	}

	var config config
	if err = json.NewDecoder(file).Decode(&config); err != nil {
		log.L.Fatal("failed to read json", zap.Error(err))
	}

	Lavalink = config.Lavalink
	Spotify = config.Spotify
	Token = config.Token
	DevGuildID = config.DevGuildID
}
