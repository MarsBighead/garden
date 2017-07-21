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
				Chromosome: gene.Chromosome,
				Strand:     gene.Strand,
				TxStart:    gene.TxStart,
				TxEnd:      gene.TxEnd,
				ExonCount:  gene.ExonCount,
				Score:      gene.Score,
				Gene:       gene.Gene,
			},
			ExonFrame: strings.Split(strings.TrimRight(string(gene.ExonFrames), ","), ","),
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
	sql = fmt.Sprintf(`select name,
							  chrom,
							  strand,
							  txStart,
							  txEnd,
							  cdsStart,
							  cdsEnd,
							  exonCount,
							  exonStarts,
							  exonEnds,
                              score,
							  name2 gene,
							  exonFrames 
					   from hg38.refGene
                       %s`, where)
	//limit 1,5`, where)
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
