package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type A struct {
	X int
	Y bool
	C string
}

var s1 = rand.NewSource(time.Now().UnixNano())
var r1 = rand.New(s1)

func make(num int) (ret []A) {
	for i := 0; i < num; i++ {
		ret = append(ret, A{X: i, C: fmt.Sprintf("id%d.1", r1.Intn(100))})
	}
	return ret
}

func main() {
	//file, err := os.OpenFile("daily_learn/year_2024/day_0402/xx.json", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//defer file.Close()
	//fmt.Println(err)
	//
	//encoder := json.NewEncoder(file)
	//for i := 0; i < 5; i++ {
	//	encoder.Encode(make(10))
	//}

	file, err := os.Open("./daily_learn/year_2024/day_0402/xx.json")
	fmt.Println(err)
	defer file.Close()

	//var list []A
	var batch []map[string]any
	decoder := json.NewDecoder(file)
	for decoder.More() {
		if err := decoder.Decode(&batch); err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(batch)
		var a A
		for _, v := range batch {
			content, _ := jsoniter.Marshal(v)
			json.Unmarshal(content, &a)
			fmt.Printf("%+v\n", a)
		}
		break
		//list = append(list, batch...)
	}
	//fmt.Println(len(list))

}
