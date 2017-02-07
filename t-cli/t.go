package client

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
        "time"
	"garden/models"
	"github.com/golang/protobuf/proto"
)

// Crawl Send http request and get response data
func Crawl(uri string) (*http.Response, error) {
	//host, _ := os.Hostname()
	//userAgent := "m" + "/3.0" + " (Appcoach, " + host + ")"
	fmt.Printf("user url is :%s\n", uri)
	// proxyURL, _ := url.Parse("http://proxy.appcoachs.net:7070")
	client := &http.Client{
	//	Transport: &http.Transport{
	//		Proxy: http.ProxyURL(proxyURL),
	//	},
		Timeout: time.Duration(300 * time.Second),
	}

	//resp, err := client.Get(uri)
	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-protobuffer")
	resp, err := client.Do(req)
	return resp, nil
}

// ReadResponse Get body data from http response
func ReadResponse(resp *http.Response) ([]byte, error) {
	fmt.Printf("Read values!\n")
	if resp == nil {
		return nil, errors.New("response is nil")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Body values: %v\n", body)
	newTest := &models.Test{}
	err = proto.Unmarshal(body, newTest)
	fmt.Printf("Unmarshall protobuf data is:\n%v\n", newTest)
	body = bytes.Replace(body, []byte("\x00"), []byte("\x20"), -1)
	fmt.Printf("Protobuf data with blank/special charater replaced is:\n%v\n", newTest)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode: %v, %s", resp.StatusCode, string(body))
	}
	return body, nil
}
