package api 

import (
    "log"

    "net/http"
    "fmt"

    "garden/models"
    "github.com/golang/protobuf/proto"
)

func GetPb(writer http.ResponseWriter) {
    test := &models.Test{
        Label: proto.String("hello"),
        Type:  proto.Int32(18),
        Reps:  []int64{1, 2, 3},
        Optionalgroup: &models.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }
    data, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    fmt.Printf("Data marshalled with protobuf:\n%v\n", data)
    newTest := &models.Test2{}
    err = proto.Unmarshal(data, newTest)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    fmt.Fprintf(writer, string(data))
    fmt.Printf("Unmashalled protobuf:\n%v,\nMarshalled data: %v\n", newTest, data)
    fmt.Printf("Independent item out put:\n    Label:%s\n", *newTest.Label)
    fmt.Printf("    Reps:%v\n", newTest.Reps)
    fmt.Printf("    OptionalGroup:%v\n", *newTest.Optionalgroup.RequiredField)
    // Now test and newTest contain the same data.
    if test.GetLabel() != newTest.GetLabel() {
        log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
    }
}
