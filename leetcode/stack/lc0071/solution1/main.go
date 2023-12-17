package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func simplifyPath1(path string) string {
	return filepath.Clean(path)
}

const (
	slash int = iota
	onepoint
	twopoint
	dir
)

func simplifyPath(path string) string {
	state := slash
	word := ""
	directories := make([]string, 0)
	for i := 0; i < len(path); i++ {
		switch state {
		case slash:
			if path[i] == '.' {
				word += string(path[i])
				state = onepoint
			} else if path[i] != '/' {
				word += string(path[i])
				state = dir
			}
		case onepoint:
			if path[i] == '.' {
				word += string(path[i])
				state = twopoint
			} else if path[i] != '/' {
				word += string(path[i])
				state = dir
			} else {
				word = ""
				state = slash
			}
		case twopoint:
			if path[i] == '/' {
				if len(directories) > 0 {
					directories = directories[:len(directories)-1]
				}
				word = ""
				state = slash
			} else {
				word += string(path[i])
				state = dir
			}
		case dir:
			if path[i] == '/' {
				directories = append(directories, word)
				word = ""
				state = slash
			} else {
				word += string(path[i])
				state = dir
			}
		}
	}
	switch state {
	case dir:
		directories = append(directories, word)
	case twopoint:
		if len(directories) > 0 {
			directories = directories[:len(directories)-1]
		}
	}
	return "/" + strings.Join(directories, "/")
}

func main() {
	path := "/..."
	fmt.Println(simplifyPath(path))

	path = ".../..."
	fmt.Println(simplifyPath(path))

	path = "/a/./b/../../c/"
	fmt.Println(simplifyPath(path))
}
