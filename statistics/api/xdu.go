package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"

	"garden/internal/xiaodu"
	"github.com/golang/protobuf/proto"
)

// Xdu Mock Xiaodu-Adx API request management
func Xdu(w http.ResponseWriter, req *http.Request) {
	raw, err := httputil.DumpRequest(req, true)
	checkError(err)
	//buf, err := ioutil.ReadAll(req.Body)
	fmt.Printf("All http data in raw:\n%v\n", string(raw))
	//var bidReq xiaodu.BidRequest
	bodyBuf :=getXduRequestBuf(req)
	bidReqBody := &xiaodu.BidRequest{}
	err = proto.Unmarshal(bodyBuf, bidReqBody)
	fmt.Printf("Device:\n %v\n", bidReqBody.Device)
	//req.Header.Set("Content-Type", "application/x-protobuffer")
	// Response binary data by func write from http.ResponseWriter
    resp:=getResponseBuf(bidReqBody.GetDevice())
	w.Write(resp)

}

func getXduRequestBuf( req *http.Request)[]byte{
    var requestSizeLimit int64 = 16 * 1024
	//var bidReq xiaodu.BidRequest
	limitReader := io.LimitReader(req.Body, requestSizeLimit)
	bodyBuf, err := ioutil.ReadAll(limitReader)
	checkError(err)
    fmt.Printf("Limit:\n %v\n", string(bodyBuf))
    return bodyBuf
}

func getResponseBuf(data *xiaodu.BidRequest_Device)[]byte{
    respBuf, err := proto.Marshal(data)
	//n, _ := io.Copy(os.Stdout, limitR)
	fmt.Printf("\nResp pb data:\n%v\n", string(respBuf))
	//err = proto.Unmarshal(limitR, newProtoTest)
	checkError(err)
    return respBuf
}
// checkError -Simplify error return checking
func checkError(err error) {
     if err != nil {
         fmt.Println("Fatal error ", err.Error())
      os.Exit(1)
   }
}

