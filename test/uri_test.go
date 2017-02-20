package test

import (
	"testing"
)

func TestGetMultiUri(t *testing.T) {
	strs := []string{"http://127.0.0.1:8002/pkg/os/#IsExist",
		"http://docs.ruanjiadeng.com/gopl-zh/ch1/ch1-06.html",
		"https://gobyexample.com/reading-files",
	}
	GetMultiUri(strs)
}
