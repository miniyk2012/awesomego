package week1

import (
	"slices"
)

type gradeSchool struct{}

func (gradeSchool) Multiply(a, b []byte) []byte {
	total := make([]uint8, 0)

	aline := make([]uint8, len(a)) // 乘数
	for j := len(a) - 1; j >= 0; j-- {
		aline[j] = a[j] - '0'
	}
	for i := len(b) - 1; i >= 0; i-- {
		bi := b[i] - '0'
		newLine := mulNum(aline, bi)
		total = addLine(total, newLine, len(b)-i-1)
	}
	result := make([]byte, len(total))
	for i := 0; i < len(total); i++ {
		result[i] = uint8(total[i]) + '0'
	}
	// 裁剪掉prefix zero
	var i int
	for ; i < len(result); i++ {
		if result[i] != '0' {
			break
		}
	}
	if i == len(result) {
		i = len(result) - 1
	}
	return result[i:]
}

func mulNum(aline []uint8, b uint8) (r []uint8) {
	aline = slices.Clone(aline)
	var inc uint8
	if b == 0 {
		return []uint8{0}
	}
	for i := len(aline) - 1; i >= 0; i-- {
		x := aline[i]*b + inc
		inc = x / 10
		x = x % 10
		r = append([]uint8{x}, r...)
	}
	if inc > 0 {
		r = append([]uint8{inc}, r...)
	}
	return
}

// total与往左偏移shift的newLine相加, 得到新的total
func addLine(total, newLine []uint8, shift int) (r []uint8) {
	total = slices.Clone(total)
	newLine = slices.Clone(newLine)
	paddingZero := append(newLine, make([]uint8, shift)...)
	maxLength := max(len(total), len(paddingZero))
	r = make([]uint8, maxLength)
	var inc uint8
	slices.Reverse(total)
	slices.Reverse(paddingZero)
	for i := 0; i < maxLength; i++ {
		var x uint8
		if len(total) > i && len(paddingZero) > i {
			x = total[i] + paddingZero[i] + inc
		} else if len(total) > i {
			x = total[i] + inc
		} else if len(paddingZero) > i {
			x = paddingZero[i] + inc
		}
		inc = x / 10
		x = x % 10
		r[i] = x
	}
	if inc > 0 {
		r = append(r, inc)
	}
	slices.Reverse(r)
	return
}
