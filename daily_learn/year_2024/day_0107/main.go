// heap剖析: https://www.cnblogs.com/huxianglin/p/6925119.html

package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int // 定义一个类型

func (h IntHeap) Len() int { return len(h) } // 绑定len方法,返回长度
func (h IntHeap) Less(i, j int) bool { // 绑定less方法
	return h[i] < h[j] // 如果h[i]<h[j]生成的就是小根堆，如果h[i]>h[j]生成的就是大根堆
}
func (h IntHeap) Swap(i, j int) { // 绑定swap方法，交换两个元素位置
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} { // 绑定pop方法，从最后拿出一个元素并返回
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(int))
}

type Heap struct {
	s []any
}

func (h *Heap) Push(x any) {
	h.s = append(h.s, x)
}

func main() {
	h := &IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0} // 创建slice
	heap.Init(h)                                // 初始化heap
	fmt.Println(*h)
	fmt.Println(heap.Pop(h)) // 调用pop
	heap.Push(h, 6)          // 调用push
	fmt.Println(*h)
	for len(*h) > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()

	a := []int{1, 2, 3}
	var b = &a
	old := *b
	*b = old[:2]
	old[2] = 4
	fmt.Printf("len=%d, cap=%d\n", len(*b), cap(*b))   // len=2, cap=3
	fmt.Printf("len=%d, cap=%d\n", len(a), cap(a))     // len=2, cap=3
	fmt.Printf("len=%d, cap=%d\n", len(old), cap(old)) // len=3, cap=3
	fmt.Printf("%v, %v, %v\n", b, a, old)
	*b = append(*b, 10)
	fmt.Printf("%v, %v, %v\n", b, a, old)

	h2 := Heap{}
	h2.Push(3)
	fmt.Printf("h2=%v", h2)
}
