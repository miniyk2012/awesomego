package concurrent_test

import (
	"bytes"
	"fmt"
	"sync"
	"testing"

	"github.com/miniyk2012/awesomego/datastructure/concurrent"
	"github.com/stretchr/testify/assert"
)

func TestWriteBytes(t *testing.T) {
	buf := bytes.NewBuffer(make([]byte, 1024))
	expect := 0
	for j := 0; j < 1024; j++ {
		buf.WriteByte(byte(j / 256))
		expect += j / 256
	}
	t.Logf("expect=%d", expect)
	assert.Equal(t, 1536, expect)
}

func ExampleExchanger() {
	buf1 := bytes.NewBuffer(make([]byte, 1024))
	buf2 := bytes.NewBuffer(make([]byte, 1024))

	exchanger := concurrent.NewExchanger[*bytes.Buffer]()

	var wg sync.WaitGroup
	wg.Add(2)

	expect := 0
	go func() {
		defer wg.Done()

		buf := buf1
		for i := 0; i < 10; i++ {
			for j := 0; j < 1024; j++ {
				buf.WriteByte(byte(j / 256))
				expect += j / 256
			}

			buf = exchanger.Exchange(buf)
		}
	}()

	var got int
	go func() {
		defer wg.Done()

		buf := buf2
		for i := 0; i < 10; i++ {
			buf = exchanger.Exchange(buf)
			for _, b := range buf.Bytes() {
				got += int(b)
			}
			buf.Reset()
		}
	}()

	wg.Wait()

	fmt.Println(got)
	fmt.Println(expect == got)

	// Output:
	// 15360
	// true
}
