package main

import (
	"fmt"
	"strings"
)

// 我用的是找规律法, 逻辑过于复杂了
func convert(s string, numRows int) string {
	length := len(s)
	if numRows >= length {
		return s
	}
	if numRows == 1 {
		return s
	}
	var builder strings.Builder
	for i := 0; i < numRows; i++ {
		builder.WriteByte(s[i])
		start := i
		for builder.Len() < length {
			if i != 0 && i != numRows-1 {
				if start+2*(numRows-1-i) < length {
					builder.WriteByte(s[start+2*(numRows-1-i)])
				} else {
					break
				}
			}
			if start+2*(numRows-1) < length {
				builder.WriteByte(s[start+2*(numRows-1)])
			} else {
				break
			}
			start += 2 * (numRows - 1)
		}
	}
	return builder.String()
}

// 可以使用模拟法做这件事
func convert2(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	var res = make([]strings.Builder, numRows) // 模拟构造numRows行字符串即可
	row, flag := 0, -1
	for i := 0; i < len(s); i++ {
		res[row].WriteByte(s[i])
		if row == 0 || row == numRows-1 {
			flag = -flag
		}
		row += flag
	}
	var ret strings.Builder
	for _, r := range res {
		ret.WriteString(r.String())
	}
	return ret.String()
}

func main() {
	fmt.Println(convert2("PAYPALISHIRING", 3))
	fmt.Println(convert2("PAYPALISHIRING", 4))
	fmt.Println(convert2("A", 1))
	fmt.Println(convert2("A", 2))
	fmt.Println(convert2("AB", 1))
}
