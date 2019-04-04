package transformer

import (
	"fmt"
	"testing"
)

type Human struct {
	ID    int64  `tsql:"column:id;type:int;default:NULL;primary_key"`
	Name  string `tsql:"column:Name;type:varchar(128);default:'Unknow'"`
	Age   int64  `sql:"-"`
	Score int64  `tsql:"column:score;type:real;"`
}

func TestCreateWithNoTablename(t *testing.T) {
	h := Human{}
	sql := CreateSQLWithoutTablename(h)
	fmt.Printf("sql is:\n%v\n", sql)
	//GetColumn(h)

}

func TestSStmtInsert(t *testing.T) {
	h := Human{
		ID: 1,
	}
	SStmtInsert(h)

}
