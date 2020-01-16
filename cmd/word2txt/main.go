package main

import (
	"fmt"
	"garden/pkg/word2txt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

func main() {
	//doc, err := document.Open("document.docx")
	fmt.Println(`
Welcome to use Word to Text program!
Next process with key Enter, and exit with character q, 
Error input tolerance 3 times.
	`)
	fmt.Println("OS type ", runtime.GOOS)
	var sig bool
	var n int
	var ss string
	for !sig {
		if n >= 5 {
			break
		}
		fmt.Printf(`Input your file/folder path or exit with q: `)
		fmt.Scanln(&ss)
		if ss == "" {
			continue
		} else {
			if strings.ToLower(ss) == "q" || strings.ToLower(ss) == "exit" {
				sig = true
			} else {
				if isExistPath(ss) {
					err := execute(ss)
					if err != nil {
						log.Println(err)
						n++
						continue
					}

				} else {
					fmt.Println(`Invaild file path input, system will exit in 10 seconds...`)
					select {
					case <-time.After(10 * time.Second):
						sig = true
					}
				}
			}
		}
	}
	log.Println(`Exit.`)

}

func isExistPath(path string) bool {
	_, err := os.Stat(path)
	if err != nil && err.Error() == os.ErrNotExist.Error() {
		return false
	}
	return true
}

func execute(pathname string) error {
	//c := new(word2txt.Converter)
	//c.Filename = "document.docx"
	//c.Extract()
	var isWindows bool
	if runtime.GOOS == "windows" {
		isWindows = true
	}
	stats, _ := os.Stat(pathname)
	if stats.IsDir() {
		fmt.Println(pathname)
		files, err := ioutil.ReadDir(pathname)
		if err != nil {
			return err
		}
		absPathname, _ := filepath.Abs(pathname)
		fmt.Println("abs path ", absPathname)
		for _, f := range files {
			filename := absPathname + "/" + f.Name()
			dst, ok := isWord(filename)
			if ok {
				fmt.Println(filename)
				c := &word2txt.Converter{
					Dst:       dst,
					IsWindows: isWindows,
				}
				err = c.Extract(filename)
				if err != nil {
					return err
				}
				err = c.Output()
				if err != nil {
					return err
				}

			}
		}
	} else {
		dst, ok := isWord(pathname)
		if ok {
			c := &word2txt.Converter{
				Dst:       dst,
				IsWindows: isWindows,
			}
			err := c.Extract(pathname)
			if err != nil {
				return err
			}
			err = c.Output()
			if err != nil {
				return err
			}

		}
	}

	return nil

}

func isWord(filename string) (string, bool) {
	reDoc := regexp.MustCompile("(doc)$")
	reDocx := regexp.MustCompile("(docx)$")
	ok := reDoc.MatchString(filename) || reDocx.MatchString(filename)
	var dst string
	if ok {
		if reDoc.MatchString(filename) {
			dst = strings.TrimSuffix(filename, "doc") + "txt"
		}
		if reDocx.MatchString(filename) {
			dst = strings.TrimSuffix(filename, "docx") + "txt"
		}
	}
	return dst, ok

}
