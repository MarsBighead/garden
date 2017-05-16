package client

import "testing"

func TestXiaoduRequest(t *testing.T) {
	// Notice that it is whole difference between
	// http://127.0.0.1:8001/api/pb and 127.0.0.1:8001/api/pb
	XiaoduRequest("xiaodu.bin", "http://127.0.0.1:8001/api/xiaodu")
}
