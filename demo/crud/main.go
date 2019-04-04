package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"log"

	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// Config  configure file for application
type Config struct {
	Application string `toml:"application"`
	Timeline    string `toml:"timeline"`
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
type Conversion struct {
	TID         string `db:"tid"`
	DeviceIP    string `db:"ip"`
	InstallTime string `db:"cov_created_at"`
	ClickTime   string `db:"clk_created_at"`
	Affiliate   string `db:"media_name"`
	Product     string `db:"external_id"`
	Params      string `db:"params"`
	MediaKey    string `db:"media_key"`

	AdvertiserKey string `db:"advertiser_key"`
	Partner       string `db:"advertiser_name"`
	ProductLabel  string `db:"pkgname"`
}

func isFileExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func getCovInfo(db *sqlx.DB, cfg *Config) []*Conversion {

	var condition string
	if isFileExist(cfg.Timeline) {
		created_at, _ := ioutil.ReadFile(cfg.Timeline)
		condition = "hcov.created_at>'" + string(created_at) + "'"
	} else {
		condition = "hcov.created_at>SUBTIME(now(),'00:15:00')"
	}
	fmt.Printf("condition is: %v\n", condition)
	sql := fmt.Sprintf(`
	      SELECT hcov.tid tid,
			    hcov.created_at cov_created_at,
				hcov.media_key media_key,
				hmed.media_name media_name,
				hcov.advertiser_key advertiser_key,
				hadv.advertiser_name advertiser_name,
				hclk.ip ip,
				hclk.params params,
				hclk.created_at clk_created_at,
				hcamp.external_id external_id, 
				hcamp.pkgname pkgname 
	     FROM   oro_hydra.hydra_conversion hcov,
		        oro_hydra.hydra_click hclk,
				oro_hydra.hydra_media hmed,
				oro_hydra.hydra_media hadv,
				oro_hydra.hydra_campaign hcamp 
		 WHERE  %s
				and hcov.tid=hclk.tid
				and hmed.media_key = hcov.media_key
				and hadv.media_key = hcov.advertiser_key
				and hcamp.campaign_id=hclk.campaign_id 
				and hadv.advertiser_name!=''`, condition)
	fmt.Printf("SQL is:\n%v\n", sql)

	var conversions []*Conversion
	db.Select(&conversions, sql)
	return conversions
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
