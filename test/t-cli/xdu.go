package client

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"fmt"

	
	"garden/internal/xiaodu"
	"github.com/golang/protobuf/proto"
)


//MockPbRequest Mock a request with protobuf data
func MockXduRequest(inputPb, uri string) {
	pbData, err := ioutil.ReadFile(inputPb)
	checkError(err)
	//fmt.Printf("protobuf byte:\n%v\n", pbData)
	body := &xiaodu.BidRequest{}
	err = proto.Unmarshal(pbData, body)
	checkError(err)
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	//resp, err := client.Post(uri, string(pbData), nil)
	reqBody := bytes.NewReader(pbData)
	req, err := http.NewRequest("POST", uri, reqBody)
	checkError(err)
	req.Header.Set("Content-Type", "application/x-protobuffer")
	fmt.Printf("Start POST protobuf byte!\n")
	resp, err := client.Do(req)
	var responseSizeLimit int64 = 16 * 1024
	limitResp := io.LimitReader(resp.Body, responseSizeLimit)
	buf, err := ioutil.ReadAll(limitResp)
	checkError(err)
	data := &xiaodu.BidRequest_Device{}
	err = proto.Unmarshal(buf, data)
	fmt.Printf("%v", string(buf))
	jOut,err := os.Create("resp.bin")
	jOut.Write(buf)
   	jOut.Close()
}

// checkError -Simplify error return checking
func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
