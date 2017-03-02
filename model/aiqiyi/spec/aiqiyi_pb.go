package spec

import (
	"io/ioutil"
	"log"
	"os"

	"appcoachs.net/x/aiqiyi"

	"encoding/json"

	"fmt"

	"github.com/golang/protobuf/proto"
)

// GenRequest Generate Xiaodu-adx request protobuf file from JSON
func genRequest(input, outputfile string) {
	//fmt.Printf("Generate Xiaodu request protobuf data from JOSN file %s.\n", inputFile)
	req := aiqiyi.BidRequest{}
	// Read json data and assign to data
	body, err := ioutil.ReadFile(input)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// Unmarshall JSON data and binding value with struct req
	err = proto.UnmarshalText(string(body), &req)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	//fmt.Printf("Marshal protobuf!\n")
	jData, err := json.Marshal(req)
	if err != nil {
		log.Fatal("marshal json error: ", err)
	}
	fmt.Printf("%v\n", string(jData))
	pData, err := proto.Marshal(&req)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// Output protobuf file
	//fmt.Printf("Write data in protobuf file %s\n", outputfile)
	fout, err := os.Create(outputfile)
	fout.Write(pData)
	fout.Close()
}
