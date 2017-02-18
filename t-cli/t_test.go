package client

import "testing"
import "fmt"

func TestClient(t *testing.T) {
	// Notice that it is whole difference between
	// http://127.0.0.1:8001/api/pb and 127.0.0.1:8001/api/pb
	//resp, _ := Crawl("http://127.0.0.1:8080/api/pb")
	resp, _ := Crawl("http://127.0.0.1:8080/hello")
	fmt.Printf("Response status: %v\n", resp.Status)
	//ReadResponse(resp)
}
