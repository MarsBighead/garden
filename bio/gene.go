package bio

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/schema"
)

func parseQueryParams(params url.Values) (q *Query, err error) {
	q = new(Query)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err = dec.Decode(q, params); err != nil {
		log.Fatal(err)
		return
	}
	return
}

//Hg38Refgene Gene information include all cut-mode
func (s *Service) Hg38Refgene(w http.ResponseWriter, r *http.Request) {
	q, err := parseQueryParams(r.URL.Query())
	if err != nil {
		log.Panic(err)
	}

	respGenes, err := q.queryRefGene(s.DB)
	if err != nil {
		log.Panic(err)
	}
	body, err := json.MarshalIndent(respGenes, "", "    ")
	if err != nil {
		log.Panic(err)
	}
	w.Write(body)
}

func (q *Query) queryRefGene(db *sql.DB) ([]*ResponseRefgene, error) {
	var whereConds []string
	var where string
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
	if where == "" {
		where = "limit 1,5"
	}
	sql := fmt.Sprintf(`select name, chrom, strand, txStart, txEnd,
	cdsStart, cdsEnd, exonCount, exonStarts, exonEnds, 
	score, name2 gene, exonFrames 
	from hg38.refGene
	%s`, where)
	//`, where)
	var resp []*ResponseRefgene
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		refGene := new(ResponseRefgene)
		err = rows.Scan(&refGene.ModeName, &refGene.Chromosome, &refGene.Strand, &refGene.TxStart, &refGene.TxEnd,
			&refGene.CdsStart, &refGene.CdsEnd, &refGene.ExonCount, &refGene.ExonStarts, &refGene.ExonEnds,
			&refGene.Score, &refGene.Gene, &refGene.ExonFrames)
		if err != nil {
			return nil, err
		}
		refGene.ExonPos = refGene.getExonPos()
		refGene.ExonFrame = strings.Split(strings.TrimRight(string(refGene.ExonFrames), ","), ",")
		resp = append(resp, refGene)

	}
	//	resp.Count = len(resp.Modes)
	return resp, nil
}

func (r *RowRefgene) getExonPos() (exonPos []*Exon) {
	sStarts := strings.Split(strings.TrimRight(string(r.ExonStarts), ","), ",")
	sEnds := strings.Split(strings.TrimRight(string(r.ExonEnds), ","), ",")
	for i, p := range sStarts {
		exS, err := strconv.Atoi(p)
		if err != nil {
			log.Fatal(err)
			continue
		}
		exE, err := strconv.Atoi(sEnds[i])
		if err != nil {
			log.Fatal(err)
			continue

		}
		exonPos = append(exonPos,
			&Exon{
				ExonStart: exS,
				ExonEnd:   exE,
			})
	}
	return
}
