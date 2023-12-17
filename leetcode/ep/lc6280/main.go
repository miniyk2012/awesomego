package main

import (
	"fmt"
	"math"
)

func getPrimes(right int) []bool {
	var primes []bool = make([]bool, right+1)
	primes[0], primes[1] = false, false
	for i := 2; i < right+1; i++ {
		primes[i] = true
	}
	// 筛法获取素数
	for i := 2; i <= int(math.Sqrt(float64(right))); i++ {
		if !primes[i] {
			continue
		}
		for x := i * i; x <= right; x = x + i {
			primes[x] = false
		}
	}
	return primes
}

func closestPrimes(left int, right int) []int {
	var ans []int = []int{-1, -1}
	primes := getPrimes(right)
	best := math.MaxInt // 最小间隔
	// last是上一个素数的位置
	for i, last := left, -1; i <= right; i++ {
		if primes[i] {
			if last > -1 && i-last < best {
				best = i - last
				ans = []int{last, i}
			}
			last = i
		}
	}
	return ans
}

func main() {
	for i, v := range getPrimes(100) {
		if v {
			fmt.Printf("%d\t", i)
		}
	}
}
