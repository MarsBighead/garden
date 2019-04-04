package parse

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
blog: xiaorui.cc
best_authors: ["fengyun","lee","park"]
desc:
  counter: 521
  plist: [3, 4]
`

type T struct {
	Blog    string
	Authors []string `yaml:"best_authors,flow"`
	Desc    struct {
		Counter int   `yaml:"Counter"`
		Plist   []int `yaml:",flow"`
	}
}

func parseYAML() {
	t := T{}
	//把yaml形式的字符串解析成struct类型
	err := yaml.Unmarshal([]byte(data), &t)
	//修改struct里面的记录
	t.Blog = "this is Blog"
	t.Authors = append(t.Authors, "myself")
	t.Desc.Counter = 99
	fmt.Printf("--- t:\n%v\n\n", t)
	//转换成yaml字符串类型
	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))
}
