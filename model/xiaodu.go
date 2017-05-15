package model

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"

	"garden/marble/xiaodu"

	"github.com/golang/protobuf/proto"
)

// FromXiaodu  Protobuf API request management
//             Response with protobuf binary data
func FromXiaodu(w http.ResponseWriter, req *http.Request) {
	raw, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Fatal("Dump http request", err)
	}
	fmt.Printf("Raw http data:\n%v\n", string(raw))
	bodyBuf := getBody(req)
	bidReqBody := new(xiaodu.BidRequest)
	err = proto.Unmarshal(bodyBuf, bidReqBody)
	resp := getResponseBuf(bidReqBody.GetDevice())
	w.Write(resp)

}

func getBody(req *http.Request) []byte {
	var requestSizeLimit int64 = 16 * 1024
	limitReader := io.LimitReader(req.Body, requestSizeLimit)
	body, err := ioutil.ReadAll(limitReader)
	if err != nil {
		log.Fatal("Read binary to buffer", err)
	}
	return body
}

func getResponseBuf(data *xiaodu.BidRequest_Device) []byte {
	respBuf, err := proto.Marshal(data)
	//n, _ := io.Copy(os.Stdout, limitR)
	fmt.Printf("\nResp pb data:\n%v\n", string(respBuf))
	//err = proto.Unmarshal(limitR, newProtoTest)
	if err != nil {
		log.Fatal("Unmarshal Device ", err)
	}
	return respBuf
}
