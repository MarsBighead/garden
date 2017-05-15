package model

import (
	"log"
	"os"
	"path/filepath"
)

//GetCurrentDir get current directory
func GetCurrentDir() (dir string) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return
}
