package main

import (
	"fmt"
	"sort"
)

func sliceDemo() {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Println("Sort by age:", people)
	fmt.Printf("slice is sorted: %t\n", sort.SliceIsSorted(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	}))
}

func searchDemo() {
	a := []int{2, 3, 4, 200, 100, 21, 234, 56}
	x := 21
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	v := sort.Search(len(a), func(i int) bool {
		return a[i] <= x
	})
	fmt.Printf("index %d is %d in %v", v, x, a)
}

func main() {
	var ints = []int{5, 3, 2, 1, 30}
	sort.Ints(ints)
	fmt.Printf("is sorted: %t\n", sort.IntsAreSorted(ints))
	fmt.Println(ints)
	fmt.Println(sort.SearchInts(ints, 4))
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))
	fmt.Println(ints)

	fmt.Println()
	sliceDemo()

	fmt.Println()
	searchDemo()
}
