package model

import (
	"log"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

//GetCurrentDir get current directory
func GetCurrentDir() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return
}

// Config  configure file for application
type Config struct {
	Application string `toml:"application"`
	Databases   struct {
		MySQL string `toml:"mysql"`
	} `toml:"databases"`
	Directory string
}

// ReadConfig  read config information
func ReadConfig() (*Config, error) {
	dir := GetCurrentDir()
	var cfg Config

	_, err := toml.DecodeFile(dir+"/config.toml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

//Parse config file to *Config
func Parse(filename string) (*Config, error) {
	cfg := new(Config)
	_, err := toml.DecodeFile(filename, cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
