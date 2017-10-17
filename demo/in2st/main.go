package main

import (
	"fmt"
	"reflect"
)

func main() {
	tonydon := &User{"TangXiaodong", 100, "0000123"}
	object := reflect.ValueOf(tonydon)
	tp := reflect.TypeOf(tonydon)
	fmt.Println(tp.String())
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	}
	tonydon.SayHello()
	v := object.MethodByName("SayHello")
	v.Call([]reflect.Value{})
}

type User struct {
	Name string
	Age  int
	Id   string
}

func (u *User) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}
