package bio

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func (s *Service) QueryAllRefGene(bucket *gocb.Bucket) error {
	sStmt := fmt.Sprintf(`select name, chrom, strand, txStart, txEnd,
	cdsStart, cdsEnd, exonCount, exonStarts, exonEnds, 
	score, name2 gene, exonFrames 
	from hg38.refGene`)
	rows, err := s.DB.Query(sStmt)
	if err != nil {
		return err
	}
	var key int
	for rows.Next() {
		refGene := new(ResponseRefgene)
		err = rows.Scan(&refGene.ModeName, &refGene.Chromosome, &refGene.Strand, &refGene.TxStart, &refGene.TxEnd,
			&refGene.CdsStart, &refGene.CdsEnd, &refGene.ExonCount, &refGene.ExonStarts, &refGene.ExonEnds,
			&refGene.Score, &refGene.Gene, &refGene.ExonFrames)
		if err != nil {
			return err
		}
		bucket.Upsert(fmt.Sprintf("%d", key), refGene, 0)

		//fmt.Printf("%v\n", string(body))
	}
	return nil
}
