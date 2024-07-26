package main

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func main() {
	//h := json.RawMessage(`{"precomputed": true}`)
	//
	//c := struct {
	//	Header *json.RawMessage `json:"header"`
	//	Body   string           `json:"body"`
	//}{Header: &h, Body: "Hello Gophers!"}
	//
	//b, err := json.MarshalIndent(&c, "", "\t")
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//os.Stdout.Write(b)

	demo()
}

type BaseTaskDetail struct {
	Env          string `json:"env""`
	Url          string `json:"url"`
	TaskStatus   int    `json:"task_status" mapstructure:"task_status"`
	Message      string `json:"message"`
	RetryHistory string `json:"retry_history" mapstructure:"retry_history"`
}

func demo() {
	data := []byte("{\"message\":\"insertDeleteTceServiceTask error\",\"task_status\":1}")
	var m map[string]interface{}
	json.Unmarshal(data, &m)
	fmt.Printf("%+v", m)
	var taskDetail BaseTaskDetail
	mapstructure.Decode(m, &taskDetail)
	fmt.Printf("%+v", taskDetail)
}
