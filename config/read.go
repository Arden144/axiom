package config

import (
	"encoding/json"
	"log"
	"os"
)

func Read(path string) (config Config) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("failed to read config: ", err)
	}

	if err = json.NewDecoder(file).Decode(&config); err != nil {
		log.Fatal("failed to read json: ", err)
	}

	return config
}
