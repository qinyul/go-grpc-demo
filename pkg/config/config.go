package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ServerAddress string `json:"server_address"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config.json")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
