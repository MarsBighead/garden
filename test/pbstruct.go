package test

import (
	"fmt"
	"garden/marble/pbt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// PbtStruct protobuf test func
func pbStruct() {
	struct1 := &pbt.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(18),
		Reps:  []int64{},
		Optionalgroup: &pbt.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	struct2 := &pbt.Test2{
		Label: proto.String("hello"),
		Reps:  []int64{},
		Optionalgroup: &pbt.Test2_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	var i, x interface{}
	i = struct1
	x = struct2
	fmt.Printf("reflect type:%v\n", reflect.TypeOf(i))
	name := fmt.Sprintf("%v", reflect.TypeOf(x))
	name = fmt.Sprintf("%v", reflect.TypeOf(i))

	switch name {
	case "*example.Test2":
		fmt.Printf("name2: %s\n", name)
	case "*example.Test":
		fmt.Printf("name1: %s\n", name)

	}
}
