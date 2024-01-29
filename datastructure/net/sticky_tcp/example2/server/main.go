package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

func main() {
	l, err := net.Listen("tcp", ":8899")
	if err != nil {
		panic(err)
	}
	fmt.Println("listen to 8899")
	for {
		conn, err := l.Accept()
		if err != nil {
			panic(err)
		} else {
			go handleConn(conn)
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		var msgSize int32
		err := binary.Read(conn, binary.BigEndian, &msgSize)
		if err != nil {
			break
		}
		buf := make([]byte, msgSize)
		_, err = io.ReadFull(conn, buf)
		if err != nil {
			break
		}
		fmt.Printf("recv: %s, len=%d\n", string(buf), msgSize)
	}
}
