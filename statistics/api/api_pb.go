package api 

import (
    "log"
    "fmt"

    "garden/models"
    "github.com/golang/protobuf/proto"

)

func GetPb() []byte{
    data := getBufPb()
    fmt.Printf("Byte data by protobuf:\n%v\n", data)
    newTest := &models.Test{}
    setPbFromBuf(data, newTest)
    //fmt.Fprintf(w, string(data))
  
    fmt.Printf("Unmashalled protobuf:\n%v,\nMarshalled data: %v\n", newTest, data)
    fmt.Printf("Independent item out put:\n    Label:%s\n", *newTest.Label)
    fmt.Printf("    Reps:%v\n", newTest.Reps)
    fmt.Printf("OptionalGroup:%v\n", *newTest.Optionalgroup.RequiredField)
    // Now test and newTest contain the same data.
    return data
   
}


func getBufPb()[]byte{
 test := &models.Test{
        Label: proto.String("hello"),
        Type:  proto.Int32(18),
        Reps:  []int64{1, 2, 3},
        Optionalgroup: &models.Test_OptionalGroup{
            RequiredField: proto.String("good bye"),
        },
    }
    bufData, err := proto.Marshal(test)
    if err != nil {
        log.Fatal("marshaling error: ", err)
    }
    return bufData
}

func setPbFromBuf(data []byte, t *models.Test){
    err := proto.Unmarshal(data, t)
    if err != nil {
        log.Fatal("unmarshaling error: ", err)
    }
    //data,err:=json.Marshal(t)
    fmt.Printf("%v",t)
}