package object

import (
	"fmt"
	"garden/marble/pbt"
	"reflect"

	"github.com/golang/protobuf/proto"
)

// ObjProtobuf protobuf test sample, follow proto v2
func ObjProtobuf() {
	st1 := &pbt.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(18),
		Reps:  []int64{},
		Optionalgroup: &pbt.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	s := proto.MarshalTextString(st1)
	fmt.Printf("Struct1 marchal text:\n%v\n", s)
	bs, _ := proto.Marshal(st1)
	fmt.Printf("Struct1 marchal text:\n%v\n", string(bs))
	st2 := &pbt.Test2{
		Label: proto.String("hello"),
		Reps:  []int64{},
		Optionalgroup: &pbt.Test2_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	var i, x interface{}
	i = st1
	x = st2
	fmt.Printf("reflect type:%v\n", reflect.TypeOf(i))
	name := fmt.Sprintf("%v", reflect.TypeOf(x))
	name = fmt.Sprintf("%v", reflect.TypeOf(i))

	switch name {
	case "*medel.Test2":
		fmt.Printf("name2: %s\n", name)
	case "*model.Test":
		fmt.Printf("name1: %s\n", name)
	}
}
