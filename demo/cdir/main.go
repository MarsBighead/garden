package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(directory string) string {
	re := regexp.MustCompile(`([\w\/]+)\/\w+$`)
	return re.FindAllStringSubmatch(directory, -1)[0][1]
	//return substr(directory, 0, strings.LastIndex(directory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func main() {

	var str1, str2 string
	str1 = getCurrentDirectory()
	fmt.Printf("Current directory %s\n", str1)

	str2 = getParentDirectory(str1)
	fmt.Println(str2)

}
