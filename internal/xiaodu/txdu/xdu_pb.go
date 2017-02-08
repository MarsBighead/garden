package xdupb

import (
	//"bytes"
	"fmt"
	"os"
	"io/ioutil"
	"encoding/json"

	"garden/internal/xiaodu"
	"github.com/golang/protobuf/proto"
)

// XduPb2JSON Transfer Xiaodu data from protobuf to JSON format
func XduPb2JSON (input, output string) {
	pbData, err := ioutil.ReadFile(input)
	checkError(err)
	//fmt.Printf("Body values: %v\n", pbData)
	bidRequest := &xiaodu.BidRequest{}
	err = proto.Unmarshal(pbData, bidRequest)
	// fmt.Printf("Unmarshall protobuf data is:\n%v\n", bidRequest)
	// *bidRequest = bytes.Replace(*bidRequest, []byte("\x00"), []byte("\x20"), -1)
	// fmt.Printf("Protobuf data with blank/special charater replaced is:\n%v\n", newTest)
	jsonData, err := json.Marshal(bidRequest)
	// fmt.Printf("JSON Marshall byte data:\n%v\n", jsonData)
	// fmt.Printf("JSON format data is:\n%v\n", string(jsonData))
   	jOut,err := os.Create(output)
	jOut.Write(jsonData)
   	jOut.Close()
}

// checkError -Simplify error return checking
func checkError(err error) {
     if err != nil {
         fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
     }
}

