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
	Count         int            `json:"count"`
	MaxModeNumber int            `json:"max_mode_number"`
	Modes         []*ModesByGene `json:"modes"`
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

type CdsStat int

const (
	None = iota + 1
	Unk
	Incmpl
	Cmpl
)

var CdsStatName = map[int32]string{
	None:   "none",
	Unk:    "unk",
	Incmpl: "incmpl",
	Cmpl:   "cmpl",
}

var CdsStatValue = map[string]int32{
	"none":   1,
	"unk":    2,
	"incmpl": 3,
	"cmpl":   4,
}

func Enum(m map[int32]string, v int32) string {
	s, ok := m[v]
	if ok {
		return s
	}
	return strconv.Itoa(int(v))
}

// RowRefgene gene modes structure
type RowRefgene struct {
	ModeName     string `db:"name" json:"mode_name"`
	Chromosome   string `db:"chrom" json:"chromosome"`
	Strand       string `db:"strand"`
	TxStart      int    `db:"txStart"`
	TxEnd        int    `db:"txEnd"`
	CdsStart     int    `db:"cdsStart"`
	CdsEnd       int    `db:"cdsStart"`
	ExonCount    int    `db:"exonCount"`
	ExonStarts   []byte `db:"exonStarts"`
	ExonEnds     []byte `db:"exonEnds"`
	Score        int    `db:"score"`
	Gene         string `db:"gene" json:"-"`
	CdsStartStat string `db:"cdsStartStat"`
	CdsEndStat   string `db:"cdsStartStat"`
	ExonFrames   []byte `db:"exonFrames"`
}

// ResponseRefgene Gene structure
type ResponseRefgene struct {
	RowRefgene
	ExonPos []*Exon `json:"exon_postion"`
}

// Exon exon position start to end
type Exon struct {
	ExonStart int `json:"exon_start"`
	ExonEnd   int `json:"exon_end"`
}

//PayloadHg38RefGene Gene information include all cut-mode
func PayloadHg38RefGene(w http.ResponseWriter, r *http.Request) {
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
	var rowGenes []*RowRefgene
	var respGenes []*ResponseRefgene
	sql := q.queryRefgeneSQL()
	fmt.Printf("sql:\n%v\n", sql)
	udb := db.Unsafe()
	udb.Select(&rowGenes, sql)

	fmt.Printf("values:\n%#v\n", rowGenes)
	for _, gene := range rowGenes {
		v := &ResponseRefgene{
			RowRefgene: RowRefgene{
				ModeName:   gene.ModeName,
				Gene:       gene.Gene,
				Chromosome: gene.Chromosome,
			},
		}
		fmt.Printf("exon starts: %v\n", string(gene.ExonStarts))
		sStarts := strings.Split(strings.TrimRight(string(gene.ExonStarts), ","), ",")
		sEnds := strings.Split(strings.TrimRight(string(gene.ExonEnds), ","), ",")
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
			v.ExonPos = append(v.ExonPos,
				&Exon{
					ExonStart: exS,
					ExonEnd:   exE,
				})

		}
		respGenes = append(respGenes, v)

	}
	body, err := json.MarshalIndent(respGenes, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(body)
}

//PayloadHg38Modes  hg38 refGene count modes API
func PayloadHg38Modes(w http.ResponseWriter, r *http.Request) {
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
func (q *QueryRefGene) queryRefgeneSQL() (sql string) {
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
	//fmt.Println(sql)
	return
}
