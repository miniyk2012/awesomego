package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func get() {
	fmt.Println("get:")
	res, err := http.Get("http://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", body)
	for k, v := range res.Header { // 打印头信息
		fmt.Println(k, ":", v)
	}
}

func post() {
	fmt.Println("\npost:")
	data := make(map[string]string, 0)
	data["key"] = "001"
	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	reqBody := strings.NewReader(string(buf))
	res, err := http.Post("http://httpbin.org/post", "application/json", reqBody)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	fmt.Printf("%s", body)
}

// // 要想请求中携带header，只能使用NewRequest()方法。
func headerPost() {
	fmt.Println("\nheaderPost:")
	data := make(map[string]string, 0)
	data["key"] = "001"
	buf, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	reqBody := strings.NewReader(string(buf))
	request, err := http.NewRequest("POST", "http://httpbin.org/post", reqBody)
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(request) // 发起客户端请求
	if err != nil {
		log.Fatal(err)
	}
	body, _ := io.ReadAll(res.Body)
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		log.Fatalf("Response failed with status code: %d\n", res.StatusCode)
	}
	fmt.Printf("%s", body)
}
func main() {
	get()
	post()
	headerPost()
}
