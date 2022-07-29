package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/disgoorg/disgolink/lavalink"
	"github.com/disgoorg/snowflake/v2"
)

const path = "config.json"

var Config struct {
	Lavalink   lavalink.NodeConfig `json:"lavalink"`
	Token      string              `json:"token"`
	DevGuildID snowflake.ID        `json:"devGuildId"`
}

func init() {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read config: ", err)
	}

	if err = json.NewDecoder(file).Decode(&Config); err != nil {
		log.Fatal("failed to read json: ", err)
	}

	fmt.Println(Config)
}
