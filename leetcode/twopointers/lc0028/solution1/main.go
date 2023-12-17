package main

import "strings"

func strStr1(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
	start:
		cur := i
		if len(haystack)-cur < len(needle) {
			return -1
		}
		for j := 0; j < len(needle); j++ {
			if haystack[cur] != needle[j] {
				i++
				goto start
			}
			cur++
		}
		return i
	}
	return -1
}
