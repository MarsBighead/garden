package test

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

//func main() {
//	Crawl("http://192.168.199.33:8001/m/api/matrix")
//}

// Crawl Send http request and get response data
func Crawl(uri string) (*http.Response, error) {
    host, _ := os.Hostname()
    userAgent := "Cilent" + "/3.0" + " (Appcoach, " + host + ")"
    fmt.Printf("userAgent:%s\n", userAgent)
    proxyURL, _ := url.Parse("http://proxy.appcoachs.net:7070")
        client := &http.Client{
            Transport: &http.Transport{
            Proxy: http.ProxyURL(proxyURL),
        },
        Timeout: time.Duration(300 * time.Second),
    }
    req, err := http.NewRequest("GET", uri, nil)
    if err != nil{
        return nil,err
    }
    req.Header.Set("User-Agent", userAgent)
    resp, err := client.Do(req)
    //resp, err := client.Get(uri)
    if err != nil {
        return nil, err
    }
    return resp, nil
}

// ReadResponse Get body data from http response
func ReadResponse(resp *http.Response) ([]byte, error) {
	if resp == nil {
		return nil, errors.New("response is nil")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	body = bytes.Replace(body, []byte("\x00"), []byte("\x20"), -1)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("StatusCode: %v, %s", resp.StatusCode, string(body))
	}

	return body, nil
}
