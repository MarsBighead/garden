package model

import (
	"github.com/BurntSushi/toml"
)

// Config  configure file for application
type Config struct {
	Application string `toml:"application"`
	Databases   struct {
		MySQL string `toml:"mysql"`
	} `toml:"databases"`
	Directory string
}

//Parse config file to *Config
func Parse(path string) (*Config, error) {
	cfg := new(Config)
	_, err := toml.DecodeFile(path, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
