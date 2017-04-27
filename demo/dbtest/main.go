package main

import (
	"fmt"

	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Config  configure file for application
type Config struct {
	Application string `toml:"application"`
	Databases   struct {
		MySQL string `toml:"mysql"`
	} `toml:"databases"`
}

// Chr  chromosome information
type Chr struct {
	RS   int     `db:"rs_id"`
	Data float64 `db:"test"`
	MAF  string  `db:"MAF"`
}

func main() {
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	mysql, err := sqlx.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
		return
	}
	//log.Fatal("Ping ", mysql.Ping())
	chrs := getChrs(mysql)
	if len(chrs) > 0 {
		fmt.Printf("RS ID\tData\tMAF\n")
	}
	for _, chr := range chrs {
		fmt.Printf("%d\t%.1f\t%s\n", chr.RS, chr.Data, chr.MAF)

	}
}

func getChrs(db *sqlx.DB) []*Chr {
	var chrs []*Chr
	db.Select(&chrs, `SELECT * FROM  chr`)
	return chrs
}

func readConfig() (*Config, error) {
	var cfg Config

	_, err := toml.DecodeFile("config.toml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
