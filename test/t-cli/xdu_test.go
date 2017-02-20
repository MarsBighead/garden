package client

import "testing"

func TestMockXduRequest(t *testing.T) {
	// Notice that it is whole difference between
	// http://127.0.0.1:8001/api/pb and 127.0.0.1:8001/api/pb
	MockXduRequest("xiaodu.bin","http://127.0.0.1:8080/api/xdu")
}
