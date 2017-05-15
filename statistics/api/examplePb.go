package api

import (
	"fmt"
	"garden/marble/pbt"
	"log"

	"github.com/golang/protobuf/proto"
)

func GetPb() []byte {

	data := getBufPb()
	fmt.Printf("Byte data by protobuf:\n%v\n", data)
	newTest := &pbt.Test{}
	setPbFromBuf(data, newTest)
	//fmt.Fprintf(w, string(data))

	fmt.Printf("Unmashalled protobuf:%v\n", newTest)
	fmt.Printf("Marshalled data: %v\n", data)
	// Now test and newTest contain the same data.
	return data

}

func getBufPb() []byte {
	test := &pbt.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(18),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &pbt.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}
	bufData, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	return bufData
}

func setPbFromBuf(data []byte, t *pbt.Test) {
	err := proto.Unmarshal(data, t)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	//data,err:=json.Marshal(t)
	fmt.Printf("%v\n", t)
}
