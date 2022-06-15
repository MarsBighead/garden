package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

type Environment struct {
	Version     string     `json:"version" yaml:"version"`
	Description string     `json:"description" yaml:"description"`
	Directory   *directory `json:"dir" yaml:"dir"`
	Database    *database  `json:"data" yaml:"database"`
}

type directory struct {
	Data string `json:"data" yaml:"data"`
}

type database struct {
	//dirver means database type, such as mysql, postgres
	Driver         string `json:"driver"  yaml:"driver"`
	dataSourceName string
	Host           string `json:"host" yaml:"host"`
	Name           string `json:"name" yaml:"name"`
	User           string `json:"user"  yaml:"user"`
	Password       string `json:"password" yaml:"password"`
	Port           int    `json:"port"`
}

func (cfg *Config) NewEnvironment() (*Environment, error) {
	conf := cfg.configFile
	file, err := os.Open(conf) // For read access.
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	env := new(Environment)
	fileType := path.Ext(conf)
	if fileType == ".json" {
		return env.parseJsonConfig(data)

	} else if fileType == ".yaml" || fileType == ".yml" {
		return env.parseYamlConfig(data)
	}
	return nil, errors.New("Unsupport file format.")
}

func (env *Environment) parseJsonConfig(data []byte) (*Environment, error) {
	err := json.Unmarshal(data, env)
	if err != nil {
		return nil, err
	}
	return env, nil

}

func (env *Environment) parseYamlConfig(data []byte) (*Environment, error) {
	err := yaml.Unmarshal(data, env)
	if err != nil {
		return nil, err
	}
	return env, nil

}
