package gsql

import (
	"fmt"
	"testing"
)

type Human struct {
	ID   int64
	Name string `gorm:"column:Name;default:'Unknow'"`
	Age  int64  `sql:"column:age;"`
}

func TestGetColumn(t *testing.T) {
	fmt.Println("Test column object.")
	h := Human{}
	GetColumn(h)

}
