package client

import (
	"fmt"
	"testing"
)

//TestCrowl Test crawl protobuf test example
func TestCrawl(t *testing.T) {
	resp, _ := Crawl("http://127.0.0.1:8001/pbt")
	fmt.Printf("Status: %v\n", resp.StatusCode)
}
