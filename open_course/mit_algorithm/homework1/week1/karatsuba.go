package week1

import (
	"strconv"
	"strings"

	"github.com/miniyk2012/awesomego/utils/go_utils"
	"golang.org/x/exp/slices"
)

type karatsuba struct{}

func (karatsuba) Multiply(x, y []byte) []byte {
	x, y = paddingTwoPower(x, y)
	return multiTwoPower(x, y)
}

// a, b长度都是2的幂
func multiTwoPower(x, y []byte) []byte {
	if len(x) != len(y) {
		panic("x, y should have same length")
	}
	size := len(x)
	if !CheckPowerOfTwo(size) {
		panic("length is not power of 2")
	}
	if size <= 10 {
		return shortMul(x, y)
	}
	a, b := x[:size/2:size/2], x[size/2:]
	c, d := y[:size/2:size/2], y[size/2:]
	ac := multiTwoPower(a, c)
	bd := multiTwoPower(b, d)
	tmp1 := add(a, b)
	tmp2 := add(c, d)
	// 需要保证multiTwoPower的参数长度相同且是2的幂次
	tmp1, tmp2 = paddingTwoPower(tmp1, tmp2)
	tmp3 := multiTwoPower(tmp1, tmp2)
	mid := sub(tmp3, add(ac, bd))
	// ac * 10^n
	first := appendZeros(ac, size)
	// mid * 10^(n/2)
	second := appendZeros(mid, size/2)
	return add(add(first, second), bd)
}

// 这里假设int64不会溢出
func shortMul(x, y []byte) []byte {
	x1, _ := strconv.ParseInt(string(x), 10, 64)
	y1, _ := strconv.ParseInt(string(y), 10, 64)
	return []byte(strconv.FormatInt(x1*y1, 10))
}

func add(x, y []byte) []byte {
	maxLength := go_utils.Max(len(x), len(y))
	x, y = slices.Clone(x), slices.Clone(y)
	slices.Reverse(x)
	slices.Reverse(y)
	r := make([]byte, maxLength)
	var inc uint8
	for i := 0; i < maxLength; i++ {
		var v uint8
		if len(x) > i && len(y) > i {
			v = x[i] - '0' + y[i] - '0' + inc
		} else if len(x) > i {
			v = x[i] - '0' + inc
		} else {
			v = y[i] - '0' + inc
		}
		inc = v / 10
		v = v % 10
		r[i] = v + '0'
	}
	if inc > 0 {
		r = append(r, inc+'0')
	}
	slices.Reverse(r)
	return r
}

// 这里假设x>y
func sub(x, y []byte) []byte {
	var dec int8 // 0或1
	x, y = slices.Clone(x), slices.Clone(y)
	y = paddingZeros(y, len(x)-len(y)) // 补到同样长度
	r := make([]byte, len(y))
	for i := 0; i < len(y); i++ {
		var v int8
		rightXi := x[len(x)-i-1] - '0'
		rightYi := y[len(y)-i-1] - '0'
		v = int8(rightXi) - dec - int8(rightYi)
		if v < 0 {
			v += 10
			dec = 1
		} else {
			dec = 0
		}
		r[len(y)-i-1] = uint8(v) + '0'
	}
	return trimLeftZeros(r)
}

func trimLeftZeros(a []byte) []byte {
	tmp := strings.TrimLeft(string(a), "0")
	if tmp == "" {
		tmp = "0"
	}
	return []byte(tmp)
}

// 尾部补n个'0'
func appendZeros(x []byte, n int) (y []byte) {
	// 0不需要再补零
	if slices.Equal(x, []byte{'0'}) {
		return x
	}
	var padding = make([]byte, n)
	for i := range padding {
		padding[i] = '0'
	}
	y = append(x, padding...)
	return
}

func paddingZeros(c []byte, n int) (d []byte) {
	var padding = make([]byte, n)
	for i := range padding {
		padding[i] = '0'
	}
	d = append(padding, c...)
	return
}

// a,b的左边补'0'到2^k, 方便对半分割
func paddingTwoPower(a, b []byte) (x, y []byte) {
	var size = go_utils.Max(len(a), len(b))
	finalSize := 1
	for finalSize < size {
		finalSize *= 2
	}
	x = paddingZeros(a, finalSize-len(a))
	y = paddingZeros(b, finalSize-len(b))
	return
}

func CheckPowerOfTwo(n int) bool {
	// added one corner case if n is zero it will also consider as power 2
	if n == 0 {
		return false
	}
	return n&(n-1) == 0
}
