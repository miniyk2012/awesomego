package main

import (
	"fmt"
	"sort"
)

func main() {
	var ints = []int{5, 3, 4, 2, 1, 30}
	sort.Ints(ints)
	fmt.Printf("is sorted: %t\n", sort.IntsAreSorted(ints))
}
