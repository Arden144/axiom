package config

import (
	"github.com/disgoorg/disgolink/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

type Config struct {
	Lavalink   lavalink.NodeConfig `json:"lavalink"`
	Token      string              `json:"token"`
	DevGuildID snowflake.ID        `json:"devGuildId"`
}
