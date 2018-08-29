package bio

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/couchbase/gocb"
)

func (s *Service) QueryAllRefGene(bucket *gocb.Bucket, ch chan *ResponseRefgene) error {
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
		key++
		refGene.DocKey = fmt.Sprintf("%d", key)
		ch <- refGene
		//fmt.Printf("%v\n", string(body))
	}
	return nil
}

func (s *Service) WriteBucket(bucket *gocb.Bucket, num int) {
	ch := make(chan *ResponseRefgene, num)
	done := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		s.QueryAllRefGene(bucket, ch)
		close(ch)
		for i := 0; i < num; i++ {
			done <- true
		}
		close(done)
	}()
	for i := 0; i < num; i++ {
		go func() {
			var sig bool
			for !sig {

				select {
				case v, ok := <-ch:
					if ok {
						_, err := bucket.Upsert(v.DocKey, v, 0)
						if err != nil {
							log.Println(err)
						}
					}
				case sig = <-done:
				}
			}
		}()
	}
	wg.Wait()
	time.Sleep(10 * time.Second)
}
