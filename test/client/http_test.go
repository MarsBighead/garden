package client

import (
	"fmt"
	"testing"
)

//TestCrowl Test crawl protobuf test example
func TestHTTPClient(t *testing.T) {
	resp, _ := HTTPClient("http://127.0.0.1:8001/protobuf")
	fmt.Printf("Status: %v\n", resp.StatusCode)
}
