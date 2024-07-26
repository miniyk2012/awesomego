package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	v := filepath.Ext("aa.tar.gz")
	fmt.Println(v)
}
