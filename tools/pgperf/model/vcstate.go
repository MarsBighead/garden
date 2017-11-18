package model

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/lib/pq"
)

//Create Create table Host for struct Host
func Create(db *sql.DB, tablename string) {
	sStmt := cSQL()
	if _, err := db.Exec(sStmt); err != nil {
		log.Fatal(err)
	}

}

//Drop drop table public."VcState"
func Drop(db *sql.DB, tablename string) {
	sStmt := `DROP TABLE IF EXISTS public."VcState"`
	if _, err := db.Exec(sStmt); err != nil {
		log.Fatal(err)
	}
}

//VcState Checker struct
type VcState struct {
	VcID                    int
	EventCollectOn          bool
	EventCollectUpdateTime  string
	PollCollectOn           bool
	PollCollectUpdateTime   string
	NewConnectionOn         bool
	NewConnectionUpdateTime string
}

//Insert into table VcState
func Insert(db *sql.DB, tablename string, numRows int) {
	start := time.Now()
	vcss := genVcStates(numRows)
	sStmt := fmt.Sprintf(`INSERT INTO %s (
"vcId" , "eventCollectOn","eventCollectUpdateTime" ,
"pollCollectOn","pollCollectUpdateTime","newConnectionOn" ,"newConnectionUpdateTime") 
VALUES ($1, $2, $3 ,$4, $5, $6, $7)
`, pq.QuoteIdentifier(tablename))
	//fmt.Println("length of vcState", len(vcss))
	var i int
	for _, vcs := range vcss {
		//fmt.Println("Rank of", i)
		modelExec(db, &vcs, sStmt)
		i++
		//i = modelQueryRow(db, &vcs, sStmt)

	}
	end := time.Now()
	dur := end.Sub(start)
	fmt.Println("Insert values duration", dur)
}

//OnlyUpdate Only update, no select values comparing
func OnlyUpdate(db *sql.DB, tablename string, vcID int) {
	now := time.Now().Format("2006-01-02 15:04:05.999999999Z07:00")
	on := true
	sStmt := fmt.Sprintf(`UPDATE "%s" SET "eventCollectOn" = %v, "eventCollectUpdateTime" = '%v'
WHERE "vcId" =$1 
`, tablename, on, now)

	switch vcID % 3 {
	case 1:
		sStmt = fmt.Sprintf(`UPDATE "%s" SET "pollCollectOn" = %v, "pollCollectUpdateTime" = '%v'
WHERE "vcId" =$1 
`, tablename, on, now)

	case 2:
		sStmt = fmt.Sprintf(`UPDATE "%s" SET "newConnectionOn" = %v, "newConnectionUpdateTime" = '%v'
WHERE "vcId" =$1 
`, tablename, on, now)
	}
	_, err := db.Exec(sStmt, vcID)
	if err != nil {
		log.Fatal("rows number", err)
	}
}

//CmpUpdate  UPDATE values based SELECT result set from  "VcState"
func CmpUpdate(db *sql.DB, tablename string, vcID int) {
	vcs := VcState{
		VcID: vcID,
	}
	sStmt := fmt.Sprintf(`SELECT * FROM "%s"  WHERE "vcId" =$1 `, tablename)
	err := db.QueryRow(sStmt, vcID).Scan(&vcs.VcID, &vcs.EventCollectOn, &vcs.EventCollectUpdateTime,
		&vcs.PollCollectOn, &vcs.PollCollectUpdateTime, &vcs.NewConnectionOn, &vcs.NewConnectionUpdateTime)

	if err != nil {
		log.Fatal("rows number", err)
	}
	now := time.Now().Format("2006-01-02 15:04:05.999999999Z07:00")
	switch vcID % 3 {
	case 0:
		vcs.EventCollectOn = !vcs.EventCollectOn
		vcs.EventCollectUpdateTime = now
	case 1:
		vcs.PollCollectOn = !vcs.PollCollectOn
		vcs.PollCollectUpdateTime = now
	case 2:
		vcs.NewConnectionOn = !vcs.NewConnectionOn
		vcs.NewConnectionUpdateTime = now
	}

	sStmt = fmt.Sprintf(`UPDATE "%s" SET "eventCollectOn" = %v, "eventCollectUpdateTime" = '%v',
"pollCollectOn" = %v, "pollCollectUpdateTime" = '%v',
"newConnectionOn" = %v, "newConnectionUpdateTime" = '%v'
WHERE "vcId" = %v`,
		tablename, vcs.EventCollectOn, vcs.EventCollectUpdateTime,
		vcs.PollCollectOn, vcs.PollCollectUpdateTime, vcs.NewConnectionOn, vcs.NewConnectionUpdateTime,
		vcID)
	_, err = db.Exec(sStmt)
	if err != nil {
		log.Fatal("rows number", err)
	}
}

func modelQueryRow(db *sql.DB, vcs *VcState, sStmt string) int {
	var i int
	sStmt = fmt.Sprintf(`%s RETURNING "vcId"`, sStmt)
	err := db.QueryRow(sStmt, vcs.VcID, vcs.EventCollectOn, vcs.EventCollectUpdateTime,
		vcs.PollCollectOn, vcs.PollCollectUpdateTime, vcs.NewConnectionOn, vcs.NewConnectionUpdateTime,
	).Scan(&i)
	if err != nil {
		log.Fatal("rows number", err)
	}
	return i
}

func modelExec(db *sql.DB, vcs *VcState, sStmt string) {
	_, err := db.Exec(sStmt, vcs.VcID, vcs.EventCollectOn, vcs.EventCollectUpdateTime,
		vcs.PollCollectOn, vcs.PollCollectUpdateTime, vcs.NewConnectionOn, vcs.NewConnectionUpdateTime,
	)
	if err != nil {
		log.Fatal("rows number", err)
	}

}

func genVcStates(l int) []VcState {
	start := time.Now()
	vcss := make([]VcState, l)
	for i := 0; i < l; i++ {
		//vcs := genVcState(i)
		vcss[i] = genVcState(i)
	}
	end := time.Now()
	dur := end.Sub(start)
	fmt.Printf("Generate %v values duration %v\n", l, dur)
	return vcss

}

func genVcState(i int) VcState {
	now := time.Now().Format("2006-01-02 15:04:05.999999999Z07:00")
	vcs := VcState{
		VcID: i + 1,
		EventCollectUpdateTime:  now,
		EventCollectOn:          true,
		PollCollectUpdateTime:   now,
		PollCollectOn:           true,
		NewConnectionUpdateTime: now,
		NewConnectionOn:         true,
	}
	//rand.Seed(time.Now().UnixNano())
	s := rand.Intn(3)
	switch s {
	case 0:
		vcs.EventCollectOn = false
	case 1:
		vcs.PollCollectOn = false
	case 2:
		vcs.NewConnectionOn = false
	}
	return vcs
}

//cSQL create table SQL
func cSQL() string {
	return `
CREATE UNLOGGED TABLE public."VcState"(
    "vcId" integer NOT NULL,
    "eventCollectOn" boolean NOT NULL,
    "eventCollectUpdateTime" timestamp with time zone NOT NULL,
    "pollCollectOn" boolean NOT NULL,
    "pollCollectUpdateTime" timestamp with time zone NOT NULL,
    "newConnectionOn" boolean NOT NULL,
    "newConnectionUpdateTime" timestamp with time zone NOT NULL,
    CONSTRAINT "VcState_pkey" PRIMARY KEY ("vcId")
);`
}
