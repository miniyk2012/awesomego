package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

func main() {
	data := []byte("~测试数据：一二三四五~")
	conn, err := net.Dial("tcp", ":8899")
	if err != nil {
		panic(err)
	}
	for i := 0; i < 2000; i++ {
		var buf [4]byte
		bufs := buf[:]
		binary.BigEndian.PutUint32(bufs, uint32(len(data)))
		if _, err := conn.Write(bufs); err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
		if _, err = conn.Write(data); err != nil {
			fmt.Printf("write failed , err : %v\n", err)
			break
		}
	}
}
