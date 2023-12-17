package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	// dp[i][j]表示s[i..j]是否是回文串
	var dp [][]bool
	var maxLength = 1
	var start int
	for i := 0; i < len(s); i++ {
		dp = append(dp, make([]bool, len(s)))
		dp[i][i] = true
	}
	// 长度从2开始逐渐递增, 这样就能保证dp正确得被填充
	for L := 2; L <= len(s); L++ {
		for i := 0; i < len(s); i++ {
			j := i + L - 1
			if j >= len(s) {
				break
			}
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				if L <= 3 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] && L > maxLength {
				maxLength = L
				start = i
			}
		}
	}
	return s[start : start+maxLength]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
	fmt.Println(longestPalindrome("xdabadx"))
}
