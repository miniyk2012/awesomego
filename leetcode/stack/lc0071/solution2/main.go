package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	stack := make([]string, 0)
	pathList := strings.Split(path, "/")
	for _, cur := range pathList {
		if cur == ".." {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else if cur != "." && cur != "" {
			stack = append(stack, cur)
		}
	}
	return "/" + strings.Join(stack, "/")
}

func main() {
	path := "/../"
	fmt.Printf("%#v\n", strings.Split(path, "/"))
	fmt.Println(simplifyPath("/../"))
}
