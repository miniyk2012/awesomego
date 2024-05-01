package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	path := "datastructure/net/sticky_tcp/example0"
	v := filepath.Base(path)
	fmt.Printf("base path=%s", v)
}
