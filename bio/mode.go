package bio

import (
	"database/sql"
	"encoding/json"
	"log"

	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

//Hg38RefgeneMode  hg38 refGene count modes API
func (s *Service) Hg38RefgeneMode(w http.ResponseWriter, r *http.Request) {
	q, err := parseQueryParams(r.URL.Query())
	if err != nil {
		log.Panic(err)
	}
	resp, err := q.queryRefGeneMode(s.DB)
	if err != nil {
		log.Panic(err)
	}
	body, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	w.Write(body)
}

func queryModeParams(uv url.Values) (q *Query, err error) {
	q = new(Query)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err = dec.Decode(q, uv); err != nil {
		return
	}
	return
}

func (q *Query) queryRefGeneMode(db *sql.DB) (*ModesResponse, error) {
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
	if q.Gene != "" {
		whereConds = append(whereConds, "name2='"+q.Gene+"'")
	}
	if len(whereConds) >= 1 {
		where = "where " + strings.Join(whereConds, " and ")
	}
	sql := fmt.Sprintf(`select count(name) mode_number, name2 gene, chrom 
		from hg38.refGene 
        %s
		group by gene, chrom
		%s
		order by mode_number desc`, where, having)
	resp := new(ModesResponse)
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		m := new(ModesByGene)
		err = rows.Scan(&m.ModeNumber, &m.Gene, &m.Chromosome)
		if err != nil {
			return nil, err
		}
		resp.Modes = append(resp.Modes, m)
		if m.ModeNumber > resp.MaxModeNumber {
			resp.MaxModeNumber = m.ModeNumber
		}
	}
	resp.Count = len(resp.Modes)
	return resp, nil
}
