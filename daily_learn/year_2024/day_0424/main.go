package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	//v := strings.SplitN("git@code.byted.org:ee/byteview_tools_apirunner.git", "/", 2)[1]
	//extension := filepath.Ext("byteview_tools_apirunner.git")
	//v = v[:len(v)-len(extension)]
	//fmt.Printf("%s %s", v, extension)
	fmt.Println()
	filepath.WalkDir("./datastructure/net/sticky_tcp", func(path string, d fs.DirEntry, err error) error {
		fmt.Println(path, d.Name())
		if d.Name() == "example0" {
			return filepath.SkipDir
		}
		return nil
	})
}
