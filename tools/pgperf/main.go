package main

import (
	"database/sql"
	"flag"
	"fmt"
	"garden/tools/pgperf/model"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"sync"
	"time"

	_ "github.com/lib/pq"
	yaml "gopkg.in/yaml.v2"
)

var pms = []int{1, 200, 2000}

func main() {
	ps := usage()
	if len(os.Args) <= 1 {
		ps = &params{
			numConcurrency: &pms[0],
			numRows:        &pms[1],
			maxOperation:   &pms[2],
		}
		log.Println(message())
	}
	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	cfg := new(PgDatabase)
	cfg.parse(path + "/config.yaml")
	//fmt.Println("length of channel buffer", runtime.NumGoroutine())
	db, err := sql.Open(cfg.Driver, cfg.getDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	infoHw()
	//fmt.Println(time.Now(), "db number is", runtime.NumGoroutine())
	start := time.Now()

	model.Drop(db, "VcState")
	model.Create(db, "VcState")
	model.Insert(db, "VcState", *ps.numRows)
	end := time.Now()
	dur := end.Sub(start)
	fmt.Printf("CREATE And INSERT %v items total duration %v\n", *ps.numRows, dur)

	//fmt.Println("Concurrency number", *ps.numConcurrency)
	concurrencyKey := make(chan int, *ps.numConcurrency)
	var wg sync.WaitGroup
	if ps.isOnly {
		for n := 0; n < *ps.maxOperation; n++ {
			wg.Add(1)
			concurrencyKey <- n
			go func(i int) {
				defer func() {
					wg.Done()
					<-concurrencyKey
				}()
				vcID := getRandNumber(i+*ps.maxOperation, *ps.numRows)
				model.OnlyUpdate(db, "VcState", vcID)
			}(n)
		}
		wg.Wait()
		close(concurrencyKey)
		tail := time.Now()
		dur = tail.Sub(end)
	} else {
		for n := 0; n < *ps.maxOperation; n++ {
			wg.Add(1)
			concurrencyKey <- n
			go func(i int) {
				defer func() {
					wg.Done()
					<-concurrencyKey
				}()
				vcID := getRandNumber(i+*ps.maxOperation, *ps.numRows)
				model.CmpUpdate(db, "VcState", vcID)
			}(n)
		}
		wg.Wait()
		close(concurrencyKey)
		tail := time.Now()
		dur = tail.Sub(end)
	}
	avg := float64(dur.Nanoseconds()) / float64(*ps.maxOperation) / 1000
	fmt.Println("opertion avg", avg, "macroseconds")
	nps := float64(*ps.maxOperation) * (1.0 / dur.Seconds())
	fmt.Printf("UPDATE %v operations duration %v,  %v n/s  \n", *ps.maxOperation, time.Now().Sub(end), nps)
}

//PgDatabase Database configure information
type PgDatabase struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	Switch   string `yaml:"switch"`
	Port     string `yaml:"-"`
	Sslmode  string `yaml:"sslmode"`
}

// Parse config.yaml to data struct
func (cfg *PgDatabase) parse(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		log.Fatal(err)
	}
}

func (cfg *PgDatabase) getDSN() string {
	var dsn string
	if strings.ToLower(cfg.Switch) == "on" {
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DBName, cfg.Sslmode)
	}
	return dsn
}

type params struct {
	numConcurrency *int
	numRows        *int
	maxOperation   *int
	isOnly         bool
}

var numConcurrency = flag.Int("g", pms[0], "UPDATE goroutine jobs")
var numRows = flag.Int("r", pms[1], "Table ROWs number")
var maxOperation = flag.Int("op", pms[2], "Max UPDATE operation numbers")
var isOnly = flag.String("oy", "false", "Operation ONLY update or not")

func usage() *params {
	flag.Parse()
	ps := new(params)
	ps.numConcurrency = numConcurrency
	ps.numRows = numRows
	ps.maxOperation = maxOperation
	*isOnly = strings.ToLower(*isOnly)
	oyRe := regexp.MustCompile("^y|yes|t|true$")
	ps.isOnly = oyRe.MatchString(*isOnly)
	return ps
}

func message() string {
	return fmt.Sprintf(`Now %s is running with default parameters as follows
 1.Parameter -g |--g : UPDATE goroutine jobs default 
 2.Parameter -r |--r : Table ROWs number default 200
 3.Parameter -op|--op: Max UPDATE operation numbers default 2000
 4.Parameter -oy|--oy: Operation ONLY update or not, 
                       accept [Yy]|[Yy][Ee][Ss]|[Tt]|[Tt][Rr][Uu][Ee] as true, default false

More manaul with %s -h|--help 
`, os.Args[0], os.Args[0])

}

//infoHw Show hardware information
func infoHw() {
	model.InfoCPU()
	model.InfoMem()
}

func getRandNumber(scope, numRows int) int {
	//rand.Seed(time.Now().UnixNano())
	s := rand.Intn(scope)
	return s%numRows + 1
}
