package concurrence

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if len(os.Args) < 3 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		fmt.Printf("usage:%s<regexp> <files>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}
	if lineRx, err := regexp.Compile(os.Args[1]); err != nil {
		log.Fatal("invalid regexp:%s\n", err)
	} else {
		grep(lineRx, commandLineFiles(os.Args[2:]))
	}
}

func source(files []string) <-chan string {
	out := make(chan string, 1000)
	go func() {
		for _, filename := range files {
			out <- filename
		}
		close(out)
	}()
	return out
}

func filterSuffixes(suffixes []string, in <-chan string) <-chan string {
	out := make(chan string, cap(in))
	go func() {
		for filename := range in {
			if len(suffixes) == 0 {
				out <- filename
				continue
			}
			ext := strings.ToLower(filepath.Ext(filename))
			for _, suffix := range suffixes {
				if ext == suffix {
					out <- filename
					break
				}
			}
		}
		close(out)
	}()
	return out
}
