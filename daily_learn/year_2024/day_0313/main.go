package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type ConfigInfo struct {
	ConfigName string  `json:"configName"`
	configID   int64   `json:"configID"`
	Unit       string  `json:"unit"`
	Env        string  `json:"env"`
	Brand      string  `json:"brand"`
	MultiGeo   string  `json:"multiGeo"`
	Value      *string `json:"value"`
}

const text = `[
    {
        "configName": "xxx",
        "configID": 100,
        "unit": "cn",
        "env": "online",
        "brand": "feishu",
        "multiGeo": "cn|us",
        "value": "abc"
    },
    {
        "configName": "xxx2",
        "configID": 123,
        "unit": "cn",
        "env": "online",
        "brand": "feishu",
        "multiGeo": "cn|us",
        "value": null
    }
]`

type V struct {
	empty bool
	id    int
}

func sort1() {
	var values = []V{
		{true, 1},
		{false, 2},
		{true, 3},
		{false, 4},
		{true, 5},
		{false, 6},
		{true, 7},
		{true, 8},
		{true, 9},
	}
	pvalues := &values
	sort.Slice(*pvalues, func(i, j int) bool {
		return (*pvalues)[i].empty == true
	})
	fmt.Println(values)
}

func json1() {
	// 定义一个包含 nil 值的 map
	data := map[string]interface{}{
		"key": nil,
	}

	// 将 map 编码为 JSON 字符串
	jsonData, _ := json.Marshal(data)
	fmt.Printf("JSON data: %s\n", jsonData)

	// 将 JSON 字符串解码回 map
	var newData map[string]interface{}
	json.Unmarshal(jsonData, &newData)

	// 检查值是否为 nil
	if newData["key"] == nil {
		fmt.Println("Value is nil")
	} else {
		fmt.Println("Value is not nil")
	}
}

type MyStruct struct {
	Field *string `json:"field"`
}

func json2() {
	// Create a struct with a nil *string field
	myStruct := MyStruct{Field: nil}

	// Marshal the struct to a JSON string
	jsonData, _ := json.Marshal(myStruct)

	fmt.Printf("JSON data: %s\n", jsonData)

	// Unmarshal the JSON string back to a struct
	var newStruct MyStruct
	json.Unmarshal(jsonData, &newStruct)

	// Check if the *string field is nil
	if newStruct.Field == nil {
		fmt.Println("Field is nil")
	} else {
		fmt.Println("Field is not nil, value:", *newStruct.Field)
	}
}

func main() {
	//var configInfos []ConfigInfo
	//json.Unmarshal([]byte(text), &configInfos)
	//spew.Printf("%v\n", configInfos)
	//
	//v := strings.Replace("{ }", " ", "", -1)
	//fmt.Println(v)
	//v = strings.Replace("[ ]", " ", "", -1)
	//fmt.Println(v)
	//sort1()
	//json1()
	json2()
}
