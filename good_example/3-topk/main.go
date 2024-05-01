// 一道TopK的题目: https://mp.weixin.qq.com/s/qDFM-nVo-jeh9VdcBfMreA

package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

// MinHeap 定义一个最小堆
type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
	fmt.Printf("after push cap=%d, len=%d\n", cap(*h), len(*h)) // cap=16, len=10
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	fmt.Printf("after pop cap=%d, len=%d\n", cap(*h), len(*h)) // cap=16, len=9
	return x
}

func findKthLargest(nums []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}
	for i := k; i < len(nums); i++ {
		//if nums[i] > (*h)[0] {
		//	heap.Pop(h)
		//	heap.Push(h, nums[i])
		//}
		heap.Push(h, nums[i]) // 和注释掉的是等价的
		if len(*h) > k {
			heap.Pop(h)
		}
	}
	fmt.Println(*h)
	return (*h)[0]
}

func main() {
	var length = 1000
	var nums = make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = i
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(nums), func(i, j int) { nums[i], nums[j] = nums[j], nums[i] })
	v := findKthLargest(nums, 10)
	fmt.Println(v)
}
