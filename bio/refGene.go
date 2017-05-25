package bio

import (
	"encoding/json"
	"log"
	"net/http"

	"garden/model"

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
	Directory string
}

// ModesByGene Count cut style mode group by gene
// select count(name) cutmode, name2 gene from refGene group by gene having cutmode>7;
type ModesByGene struct {
	ModeNumber int    `db:"mode_number" json:"mode_number"`
	Gene       string `db:"gene"        json:"gene"`
	Chromosome string `db:"chrom"       json:"chromosome"`
}

// JSONRefGene hg38 refGene API
func JSONRefGene(w http.ResponseWriter, r *http.Request) {
	cfg, err := readConfig(model.GetCurrentDir())
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	var mbgs []*ModesByGene
	db.Select(&mbgs, `select count(name) mode_number, name2 gene, chrom from refGene group by gene,chrom`)
	body, err := json.MarshalIndent(mbgs, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func readConfig(dir string) (*Config, error) {
	var cfg Config

	_, err := toml.DecodeFile(dir+"/config.toml", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
