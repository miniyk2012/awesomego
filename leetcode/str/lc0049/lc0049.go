package main

import "sort"

type ByLetter []byte

func (b ByLetter) Len() int {
	return len(b)
}

func (b ByLetter) Less(i, j int) bool {
	return b[i] < b[j]
}

func (b ByLetter) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func groupAnagrams(strs []string) (ret [][]string) {
	group := make(map[string][]string)
	for _, str := range strs {
		l := ByLetter(str)
		sort.Sort(l)
		group[string(l)] = append(group[string(l)], str)
	}
	for _, ana := range group {
		ret = append(ret, ana)
	}
	return
}
