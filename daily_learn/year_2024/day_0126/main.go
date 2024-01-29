package main

import (
	"fmt"
)

func main() {
	c := yield()
	for val := range c {
		fmt.Printf("out: %s\n", val)
		//c <- ""
	}
}

func yield() chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		test := []string{"hello", "world", "1", "2", "3", "4"}
		for _, s := range test {
			fmt.Printf("in: %s\n", s)
			c <- s
			//<-c
		}
	}()

	return c
}
