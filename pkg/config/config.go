package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Allow_IP []string `json:"allow-ip"`
}

func ReadConfigFile() *Config {
	file, err := os.ReadFile("config.json")
	if err != nil {
		return &Config{}
	}
	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return &Config{}
	}
	return &config
}
