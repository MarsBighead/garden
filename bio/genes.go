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

func queryRefgeneParams(uv url.Values) (q *Query, err error) {
	q = new(Query)
	dec := schema.NewDecoder()
	dec.IgnoreUnknownKeys(true)
	if err = dec.Decode(q, uv); err != nil {
		log.Fatal(err)
		return
	}
	return
}

//Hg38Refgene Gene information include all cut-mode
func Hg38Refgene(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	q, err := queryRefgeneParams(params)
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
	var rowGenes []*RowRefgene
	var respGenes []*ResponseRefgene
	sql := q.refGeneSQL()
	udb := db.Unsafe()
	udb.Select(&rowGenes, sql)

	for _, gene := range rowGenes {
		v := &ResponseRefgene{
			RowRefgene: RowRefgene{
				ModeName:   gene.ModeName,
				Gene:       gene.Gene,
				Chromosome: gene.Chromosome,
			},
		}
		v.ExonPos = gene.getExonPos()
		respGenes = append(respGenes, v)
	}
	body, err := json.MarshalIndent(respGenes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

func (q *Query) refGeneSQL() (sql string) {
	sql = fmt.Sprintf(`select name,
							  name2 gene,
							  cdsStart,
							  cdsEnd,
							  exonStarts,
							  exonEnds,
							  chrom
					   from hg38.refGene
					   limit 1,5`)
	return
}

func (g *RowRefgene) getExonPos() (exonPos []*Exon) {
	sStarts := strings.Split(strings.TrimRight(string(g.ExonStarts), ","), ",")
	sEnds := strings.Split(strings.TrimRight(string(g.ExonEnds), ","), ",")
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
