package main

import (
	"fmt"
	"sort"
	"testing"
)

func TestSortByLetter(t *testing.T) {
	arrays := []string{"bat", "abt"}
	for _, str := range arrays {
		l := ByLetter(str)
		sort.Sort(l)
		fmt.Printf("%s\t", l)
	}
}

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	ret := groupAnagrams(strs)
	fmt.Println(ret)
}
