package bind

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lib/pq"
)

func Bind(v interface{}) {
	t := reflect.TypeOf(v)
	var s reflect.Value
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		s = reflect.ValueOf(v).Elem()
	}
	//	t := T{12, "someone-life", nil, nil}
	currentLocation, _ := time.LoadLocation("Local")
	now, _ := pq.ParseTimestamp(currentLocation, "2017-10-25 14:50:00.802199+08")
	//now := reflect.ValueOf().Interface()
	var isTime bool
	for i := 0; i < t.NumField(); i++ {
		//fs := t.Field(i)
		fieldValue := reflect.New(t.Field(i).Type).Interface()
		//if fs.Type
		if t.Field(i).Type.Kind() == reflect.Ptr {
			fieldValue = reflect.New(t.Field(i).Type.Elem()).Interface()
		}
		if _, ok := fieldValue.(*time.Time); ok {
			isTime = true
		}
		if isTime {
			if t.Field(i).Type.Kind() == reflect.Ptr {
				s.Field(i).Set(reflect.ValueOf(&now))
			} else {
				s.Field(i).Set(reflect.ValueOf(now))
			}
		}
	}

	s.Field(0).SetInt(123)                        // 内置常用类型的设值方法
	sliceValue := reflect.ValueOf([]int{1, 2, 3}) // 这里将slice转成reflect.Value类型
	s.FieldByName("Children").Set(sliceValue)
	fmt.Println(s)

}
