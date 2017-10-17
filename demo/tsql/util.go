package gsql

import (
	"fmt"
	"go/ast"
	"reflect"
	"strings"
)

//CreateSQLWithoutTablename Create sql from struct without tablename parameters
func CreateSQLWithoutTablename(v interface{}) string {
	tablename := reflect.ValueOf(v).Type().Name()
	fmt.Println("table name is", tablename)
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
	if _, ok := tagAttr["DEFAUL"]; ok {
		fieldSQL.Default = tagAttr["DEFAUL"]
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
