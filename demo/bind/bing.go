package bind

import (
	"fmt"
	"reflect"
	"time"

	"github.com/lib/pq"
)

func SliceBind(value interface{}) {
	v1 := reflect.ValueOf("paul1")
	v2 := reflect.ValueOf("Duan")
	t := reflect.TypeOf(value)
	val := reflect.ValueOf(value).Elem()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		switch v := val; v.Kind() {
		case reflect.Slice:
			n := v.Len()
			v.Set(reflect.Append(val, reflect.Zero(v.Type().Elem())))
			fmt.Println("Input slice address", n)
			if n == 1 {
				v.Index(n).Set(v1)
			} else if n == 2 {
				v.Index(n).Set(v2)
				return
			}
			SliceBind(value)
		}
	}

}

func Bind(v interface{}) {
	t := reflect.TypeOf(v)
	var s reflect.Value
	switch t.Kind() {
	case reflect.Ptr:
		t = t.Elem()
		s = reflect.ValueOf(v).Elem()
	case reflect.Struct:
		return
	}
	//s = reflect.ValueOf(v).Elem()
	//fmt.Println("Bind", s)
	//return
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

	s.Field(0).SetInt(123)
	sliceValue := reflect.ValueOf([]int{1, 2, 3})
	s.FieldByName("Children").Set(sliceValue)
	fmt.Println(s)
	return
}
