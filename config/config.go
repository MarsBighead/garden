package config

import (
	"context"
	"database/sql"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Version     string `json:"version" yaml:"version"`
	Description string `json:"description" yaml:"description"`
	configFile  string
	flagSet     *flag.FlagSet
	db          []Database
}

type Database struct {
	XMLName  xml.Name `xml:"databases"`
	DBType   string   `xml:"dbType"`
	DBName   string   `xml:"dbName"`
	DBUser   string   `xml:"dbUser"`
	Password string   `xml:"password"`
}

func NewConfig(ctx context.Context) *Config {
	cfg := &Config{}
	cfg.flagSet = flag.NewFlagSet("garden", flag.ContinueOnError)
	fs := cfg.flagSet
	fs.StringVar(&cfg.configFile, "config", "config.yml", "config file, support format YAML, Json")
	return cfg
}

func GetConfig(dir string) Config {
	configXML := dir + "/config/config.xml"
	file, err := os.Open(configXML) // For read access.

	if err != nil {
		fmt.Printf("error: %v", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	conf := Config{}
	err = xml.Unmarshal(data, &conf)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	return conf
}

// ConnectMySQL connect mysql databases
func ConnectMySQL(dir string) (*sql.DB, error) {
	conf := GetConfig(dir).db[0]
	dbType, dbUser, dbName, Password := conf.DBType, conf.DBUser, conf.DBName, conf.Password

	db, err := sql.Open(dbType, dbUser+":"+Password+"@/"+dbName)
	if err != nil {
		log.Fatal("Error with MySQL connect: ", err)
		return nil, err
	}
	return db, nil
}
