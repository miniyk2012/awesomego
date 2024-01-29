package main

import "fmt"

func main() {
	var message = "a"
	defer func(m string) {
		fmt.Println(m)
	}(message)
	message = "b"
}
