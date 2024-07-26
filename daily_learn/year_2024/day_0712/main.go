package main

import (
	"errors"
	"fmt"
)

func main() {
	err := panic_error()
	fmt.Printf("%v\n", err)
	fmt.Println("laste")
}

func panic_error() (err error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			err = errors.New(fmt.Sprintf("%v", r))
		}
	}()
	panic("wakao")
	return nil
}
