package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	IsDebug bool `json:"is_debug"`
	Listen  struct {
		Protocol string `json:"protocol"`
		BindIp   string `json:"bind_ip"`
		Port     string `json:"port"`
	} `json:"listen"`
}

func LoadConfiguration(file string) (Config, error) {
	var cfg Config
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return cfg, err
	}

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
