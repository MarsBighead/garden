package api 

import (
    "log"

    "net/http"
    "fmt"

    "garden/statistics/apipb"
    "github.com/golang/protobuf/proto"
)

func GetPb(writer http.ResponseWriter) {
    test := &apipb.Test{
        Label: proto.String("hello"),
        Type:  proto.Int32(18),
        Reps:  []int64{1, 2, 3},
        Optionalgroup: &apipb.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    fmt.Printf("protobuf marshal sample: %v\n", data)
    newTest := &apipb.Test2{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    fmt.Fprintf(writer, string(data))
    fmt.Printf("protobuf:\n%v, data: %v\n", newTest, data)
    fmt.Printf("Output:\n%s,%d, %v\n", *newTest.Label, newTest.Reps, *newTest.Optionalgroup.RequiredField)
    // Now test and newTest contain the same data.
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
    // etc.
}
