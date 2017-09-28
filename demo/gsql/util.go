package gsql

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

func GetColumn(v interface{}) {
	//fmt.Println("vim-go")
	rtp := reflect.ValueOf(v).Type()
	fmt.Printf("%#v\n", rtp)
	fmt.Printf("num field is %#v\n", rtp.NumField())
	fmt.Printf("kind: %v; slice %v, Ptr %v\n", rtp.Kind(), reflect.Slice, reflect.Ptr)
	for rtp.Kind() == reflect.Slice || rtp.Kind() == reflect.Ptr {
		fmt.Printf("kind: %v; slice %v, Ptr %v\n", rtp.Kind(), reflect.Slice, reflect.Ptr)
		rtp = rtp.Elem()
	}
	for i := 0; i < rtp.NumField(); i++ {
		if fst := rtp.Field(i); ast.IsExported(fst.Name) {
			fmt.Printf("Name is %v\n", fst.Name)
			fmt.Printf("Tag is %v\n", fst.Tag)
			//	if len(string(fst.Tag)) > 1 {
			fmt.Printf("key-value %v\n", getTagAttribute(fst.Tag))
			//	}

		}
	}
}
func getTagAttribute(tags reflect.StructTag) map[string]string {
	var attr = map[string]string{}
	for _, str := range []string{tags.Get("sql"), tags.Get("gorm")} {
		tags := strings.Split(str, ";")
		for _, value := range tags {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))
			if len(v) >= 2 {
				attr[k] = strings.Join(v[1:], ":")
			} else {
				attr[k] = k
			}
		}
	}
	return attr
}
