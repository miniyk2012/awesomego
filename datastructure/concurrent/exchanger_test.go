package concurrent

import (
	"bytes"
	"testing"
)

func TestExchanger(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	expect := 0
	for j := 0; j < 1024; j++ {
		buf.WriteByte(byte(j / 256))
		expect += j / 256
	}
	t.Logf("expect=%d", expect)
}
