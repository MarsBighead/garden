package bio

import "strconv"

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

// Query Parse query refGene data struct
type Query struct {
	Chromosome string `schema:"chrom"`
	Start      int    `schema:"start"`
	End        int    `schema:"end"`
	MinModes   int    `schema:"min_modes"`
	MaxModes   int    `schema:"max_modes"`
}

// CdsStat type of coding region enum
type CdsStat int

// CdsStat all
const (
	None = iota + 1
	Unk
	Incmpl
	Cmpl
)

// CdsStatName coding region stats name
var CdsStatName = map[int32]string{
	None:   "none",
	Unk:    "unk",
	Incmpl: "incmpl",
	Cmpl:   "cmpl",
}

// CdsStatValue coding region stats
var CdsStatValue = map[string]int32{
	"none":   1,
	"unk":    2,
	"incmpl": 3,
	"cmpl":   4,
}

// Enum Build enum type
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
