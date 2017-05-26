package bio

import (
	"encoding/json"
	"garden/model"
	"log"

	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/schema"
	"github.com/jmoiron/sqlx"
)

//Hg38RefgeneModes  hg38 refGene count modes API
func Hg38RefgeneModes(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := queryModeParams(params)
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
	sql := q.geneModeSQL()
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

func queryModeParams(uv url.Values) (q *Query, err error) {
	q = new(Query)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err = dec.Decode(q, uv); err != nil {
		log.Fatal(err)
		return
	}
	return
}

func (q *Query) geneModeSQL() (sql string) {
	var havingConds, whereConds []string
	if q.Chromosome != "" {
		havingConds = append(havingConds, "chrom='"+q.Chromosome+"'")
	}
	if q.MinModes != 0 {
		havingConds = append(havingConds, "mode_number>="+strconv.Itoa(q.MinModes))
	}
	if q.MaxModes != 0 {
		havingConds = append(havingConds, "mode_number<="+strconv.Itoa(q.MaxModes))
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
	return
}
