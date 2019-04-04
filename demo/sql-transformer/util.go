package transformer

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

//CreateSQLWithoutTablename Create sql from struct without parameter tablename
func CreateSQLWithoutTablename(v interface{}) string {
	tablename := reflect.ValueOf(v).Type().Name()
	fmt.Println("table name is", tablename)
	return create(v, tablename)
}

//CreateSQLWithTablename Create sql from struct with parameter tablename
func CreateSQLWithTablename(v interface{}, tablename string) string {
	return create(v, tablename)
}

//FieldOfSQL Data model for build SQL statement
type FieldOfSQL struct {
	Column       string
	Type         string
	Default      string
	IsPrimaryKey bool
}

//ModelOfSQL Data model for build SQL statement
type ModelOfSQL struct {
	Fields           []*FieldOfSQL
	Columns          []string
	CreateItems      []string
	defaultTableName string
}

func create(v interface{}, tablename string) string {
	rtp := reflect.ValueOf(v).Type()
	for rtp.Kind() == reflect.Slice || rtp.Kind() == reflect.Ptr {
		rtp = rtp.Elem()
	}
	moSQL := new(ModelOfSQL)
	moSQL.defaultTableName = tablename
	for i := 0; i < rtp.NumField(); i++ {
		if stField := rtp.Field(i); ast.IsExported(stField.Name) {
			fieldSQL := &FieldOfSQL{
				Column: stField.Name,
			}
			if len(string(stField.Tag)) > 3 {
				tagAttr := getTagAttribute(stField.Tag)
				if isSkipField(tagAttr) {
					continue
				}
				fieldSQL.attachTags(tagAttr)
			} else {
				continue
			}
			moSQL.Fields = append(moSQL.Fields, fieldSQL)
			moSQL.Columns = append(moSQL.Columns, fieldSQL.Column)
			createItem := fmt.Sprintf("%s %s %s", fieldSQL.Column, fieldSQL.Type, fieldSQL.Default)
			moSQL.CreateItems = append(moSQL.CreateItems, createItem)
		}
	}
	fmt.Println(strings.Join(moSQL.CreateItems, ",\n"))
	sql := fmt.Sprintf("CREATE TABLE %v (%v)", moSQL.defaultTableName, strings.Join(moSQL.CreateItems, ",\n"))
	return sql
}

func isSkipField(tagAttr map[string]string) bool {
	if _, ok := tagAttr["-"]; ok {
		return true
	}
	return false
}

func (fieldSQL *FieldOfSQL) attachTags(tagAttr map[string]string) {
	if _, ok := tagAttr["COLUMN"]; ok {
		fieldSQL.Column = tagAttr["COLUMN"]
	}
	if _, ok := tagAttr["TYPE"]; ok {
		fieldSQL.Type = strings.ToUpper(tagAttr["TYPE"])
	}
	if _, ok := tagAttr["DEFAULT"]; ok {
		fieldSQL.Default = tagAttr["DEFAULT"]
	}
}

func getTagAttribute(tags reflect.StructTag) map[string]string {
	var attr = map[string]string{}
	for _, str := range []string{tags.Get("sql"), tags.Get("tsql")} {
		if len(str) == 0 {
			continue
		}
		tags := strings.Split(str, ";")
		for _, value := range tags {
			v := strings.Split(value, ":")
			k := strings.TrimSpace(strings.ToUpper(v[0]))
			if len(v) >= 2 {
				attr[k] = strings.Join(v[1:], ":")
			} else if len(v) == 1 {
				attr[k] = k
			} else {
				continue
			}
		}
	}
	return attr
}

//SStmtInsert SQL Statement of insert
func SStmtInsert(o interface{}) {
	t := reflect.TypeOf(o)         //获取接口的类型
	fmt.Println("Type:", t.Name()) //t.Name() 获取接口的名称

	v := reflect.ValueOf(o) //获取接口的值类型
	fmt.Println("Fields:")

	for i := 0; i < t.NumField(); i++ { //NumField取出这个接口所有的字段数量
		f := t.Field(i)                                         //取得结构体的第i个字段
		val := v.Field(i).Interface()                           //取得字段的值
		fmt.Printf("Field %6s: %v = %v\n", f.Name, f.Type, val) //第i个字段的名称,类型,值
	}

	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type) //获取方法的名称和类型
	}
}

func isValueBlank(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
