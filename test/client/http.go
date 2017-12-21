package client

import (
	"bytes"
	"errors"
	"fmt"
	"garden/marble/pbt"
	"io/ioutil"
	"net/http"
	"time"

	"log"

	"github.com/golang/protobuf/proto"
)

// HTTPClient Send http request and get response data
func HTTPClient(uri string) (*http.Response, error) {
	client := &http.Client{
		Timeout: time.Duration(300 * time.Second),
	}
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-protobuffer")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := readResponse(resp)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n", string(body))
	fmt.Printf("Http StatusCode: %d\n", resp.StatusCode)
	return resp, nil
}

// readResponse Get body data from http response
func readResponse(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("response is nil")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pbt := new(pbt.Test)
	err = proto.Unmarshal(body, pbt)
	fmt.Printf("Protobuf  test case: %#v\n", pbt)
	body = bytes.Replace(body, []byte("\x00"), []byte("\x20"), -1)
	// fmt.Printf("Blank/special replaced: %#v\n", pbt)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode error: %d", resp.StatusCode)
	}
	return body, nil
}
