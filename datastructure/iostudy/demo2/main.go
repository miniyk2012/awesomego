package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/miniyk2012/awesomego/utils"
)

func read1() {
	var p = make([]byte, 10)
	a := p[:4]
	fmt.Printf("cap=%d, len=%d\n", cap(a), len(a))

	v := strings.NewReader("abc")
	n, err := v.Read(p)
	fmt.Printf("n=%d, err=%v\n", n, err)
	n, err = v.Read(p)
	fmt.Printf("n=%d, err=%v\n", n, err)
	fmt.Printf("err == EOF: %t", err == io.EOF)
}

func read2() {
	reader := strings.NewReader("Go语言中文网aa")
	p := make([]byte, 90)
	n, err := reader.ReadAt(p, 2)
	fmt.Printf("%s, %d %v\n", p, n, err)
}

func write1() {
	file, err := os.Create(filepath.Join(utils.GetProjectRoot(), "datastructure/iostudy/demo2/output.txt"))
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString("hello, world")
	file.WriteAt([]byte("Go语言中文网"), 7)
}
func main() {
	read1()
	fmt.Println()
	read2()
	write1()
}
