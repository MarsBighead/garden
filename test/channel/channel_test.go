package channel

import (
	"testing"
)

func TestNonBuffer(t *testing.T) {
	nonBuffer()
}

/* Test result
duanp-m01:channel duanp$ go test -v -test.run TestNonBuffer
=== RUN   TestNonBuffer
I'm waiting, but not too long!
Coffee  is ready
Fruit Juice  is ready
Tea  is ready
--- PASS: TestNonBuffer (2.00s)
PASS
ok      garden/test/channel     2.011s
duanp-m01:channel duanp$ go test -v -test.run TestNonBuffer
=== RUN   TestNonBuffer
I'm waiting, but not too long!
Coffee  is ready
Tea  is ready
Fruit Juice  is ready
--- PASS: TestNonBuffer (2.00s)
PASS
ok      garden/test/channel     2.009s
*/
