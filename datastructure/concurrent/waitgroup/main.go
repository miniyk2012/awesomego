package main

import (
	"fmt"
	"time"
)

const DayLayout = "2006-01-02 15:04:05"

func main() {
	theTime, _ := time.Parse(DayLayout, "2024-02-01 01:01:01")
	fmt.Printf("%v", theTime)
}
