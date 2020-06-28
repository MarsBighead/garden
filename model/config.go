package model

import (
	"fmt"

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

// Pointor Test data struct
type Pointor struct {
	num int
}

func (p *Pointor) get() int {
	fmt.Println("Get value")
	return p.num
}
func (p *Pointor) put(val int) {
	fmt.Println("Put value")
	p.num = val
}

// Assigner test interface usage method
type Assigner interface {
	get() int
	put(int)
}

func assign(p Assigner) {
	fmt.Println(p.get())
	p.put(1)
	fmt.Println(p.get())
}
