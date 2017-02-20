package test 

import "testing"

func TestClient(t *testing.T) {
    Crawl("http://127.0.0.1:8001/m/api/matrix")
}
