package bio

import (
	"encoding/json"
	"log"
	"net/http"

	"garden/model"

	"fmt"

	"net/url"

	"strconv"

	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
)

// ModesByGene Count cut style mode group by gene
// select count(name) cutmode, name2 gene from refGene group by gene having cutmode>7;
type ModesByGene struct {
	ModeNumber int    `db:"mode_number" json:"mode_number"`
	Gene       string `db:"gene"        json:"gene"`
	Chromosome string `db:"chrom"       json:"chromosome"`
}

// ModesResponse render modes response page
type ModesResponse struct {
	Modes         []*ModesByGene `json:"modes"`
	Count         int            `json:"count"`
	MaxModeNumber int            `json:"max_mode_number"`
}

// QueryRefGene for query data parse
type QueryRefGene struct {
	Chromosome string `schema:"chrom"       json:"chromosome"`
	Start      int    `schema:"start"`
	End        int    `schema:"end"`
	MinModes   int    `schema:"min_modes"`
	MaxModes   int    `schema:"max_modes"`
}

func getQueryRefGene(uv url.Values) (q *QueryRefGene, err error) {
	q = new(QueryRefGene)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err = dec.Decode(q, uv); err != nil {
		log.Fatal(err)
		return
	}
	return
}

//PayloadHg38GeneModes  hg38 refGene count modes API
func PayloadHg38GeneModes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := getQueryRefGene(params)
	if err != nil {
		log.Fatal(err)
	}
	cfg, err := model.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := sqlx.Open("mysql", cfg.Databases.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	var mbgs []*ModesByGene
	sql := q.geneModesSQL()
	db.Select(&mbgs, sql)
	resp := ModesResponse{
		Count:         len(mbgs),
		MaxModeNumber: mbgs[0].ModeNumber,
		Modes:         mbgs,
	}

	body, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func (q *QueryRefGene) geneModesSQL() (sql string) {
	//fmt.Printf("query parameters: %#v\n", q)
	var havingConds, whereConds []string
	if q.Chromosome != "" {
		havingConds = append(havingConds, "chrom='"+q.Chromosome+"'")
	}
	if q.MinModes != 0 {
		havingConds = append(havingConds, "mode_number>="+strconv.Itoa(q.MinModes))
	}
	if q.MaxModes != 0 {
		havingConds = append(havingConds, "mode_number>="+strconv.Itoa(q.MaxModes))
	}
	var having, where string
	if len(havingConds) >= 1 {
		having = "having " + strings.Join(havingConds, " and ")
	}
	if q.Start != 0 {
		whereConds = append(whereConds, "txStart>"+strconv.Itoa(q.Start))
	}
	if q.End != 0 {
		whereConds = append(whereConds, "txEnd>"+strconv.Itoa(q.End))
	}
	if len(whereConds) >= 1 {
		where = "where " + strings.Join(whereConds, " and ")
	}
	sql = fmt.Sprintf(` select count(name) mode_number, 
	                               name2        gene, 
							       chrom 
							from hg38.refGene 
                            %s
							group by gene,chrom
							%s
							order by mode_number desc`, where, having)
	//fmt.Println(sql)
	return
}
