package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"sync"
)

type VarInfo struct {
	VarID     int64
	VarName   string
	ValueType int32
	GenType   int32
}

type NotEmptyElement struct {
	VarInfo
	ConfigIDs []int64 // 这些config下的变量值为空
}

func JsonDemo() {
	varInfo := NotEmptyElement{VarInfo{VarName: "hello"}, []int64{1, 2, 3}}
	v, _ := json.Marshal(varInfo)
	fmt.Printf("%s", v)
	m := make(map[string]string)
	m["NotEmpty"] = string(v)
	ms, _ := json.Marshal(m)
	fmt.Printf("%s", ms)
}

func main() {
	//JsonDemo()
	//regexpDemo()
	noCopyDemo()
}

type Value struct {
	mu sync.Mutex
}

func noCopyDemo() {
	v1 := Value{}
	v2 := v1
	fmt.Println(v2)
}

func regexpDemo() {
	re, err := regexp.Compile(`(?i)key|token`)
	if err != nil {
		fmt.Printf("Error compiling regex: %v", err)
		return
	}

	// 测试字符串
	testStr := "This is an Example ks"

	// 检查是否有匹配
	matched := re.MatchString(testStr)
	fmt.Println("Matched:", matched) // 输出: Matched: true
}
