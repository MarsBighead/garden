package model

import (
	"fmt"
	"garden/marble/pbt"
	"log"
	"net/http"

	"github.com/golang/protobuf/proto"
)

// Pbt Test binary protobuf data protocol
func Pbt(w http.ResponseWriter, req *http.Request) {
	test := pbtData()
	pbt, err := proto.Marshal(test)
	if err != nil {
		log.Println("marshaling error:", err)
	}
	w.Write(pbt)
}

// RebuildPbt Test Marshal/Unmarshal protobuf
func RebuildPbt(w http.ResponseWriter, req *http.Request) {
	test := pbtData()
	data, err := proto.Marshal(test)
	if err != nil {
		log.Println("marshaling error:", err)
	}
	body := new(pbt.Test)
	err = proto.Unmarshal(data, body)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Printf("Unmarshal test to body: %#v\n", body)
	s := proto.MarshalTextString(body)
	w.Write([]byte(s))
}

func pbtData() (data proto.Message) {
	data = &pbt.Test{
		Label: proto.String("hello"),
		Type:  proto.Int32(18),
		Reps:  []int64{1, 2, 3},
		Optionalgroup: &pbt.Test_OptionalGroup{
			RequiredField: proto.String("good bye"),
		},
	}

	return
}
