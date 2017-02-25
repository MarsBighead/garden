package test

import (
	"testing"
)

func TestGetMultiUri(t *testing.T) {
	strs := []string{"http://127.0.0.1:8002/pkg/github.com/golang/protobuf/proto/",
		"https://www.zhihu.com",
		"https://gobyexample.com/reading-files",
	}
	GetMultiUri(strs)
}
