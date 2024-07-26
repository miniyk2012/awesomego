package main

import (
	"errors"
	"fmt"
)

func main() {
	//os.Setenv("HTTP_PROXY", "http://127.0.0.1:12345")
	//req, err := http.NewRequest("GET", "http://example.com", nil)
	//
	//if err != nil {
	//	panic(err)
	//}
	//url, err := http.ProxyFromEnvironment(req)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(url)
	err := wrapper()
	fmt.Printf("outer %v\n", err)
}

func wrapper() error {
	err := errors.New("bbb")
	defer func() {
		err = close()
		fmt.Printf("inner %v\n", err)
	}()
	return err
}
func close() error {
	return errors.New("aaa")
}
