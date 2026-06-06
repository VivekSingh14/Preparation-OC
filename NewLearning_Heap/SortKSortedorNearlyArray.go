package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements a max heap using container/heap
type MinHeap1 []int

func (h MinHeap1) Len() int {
	return len(h)
}

func (h MinHeap1) Less(i, j int) bool {
	return h[i] < h[j]
} // Min heap: smaller elements have priority

func (h MinHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap1) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKSmallest1(arr []int, k int) ([]int, bool) {
	if k <= 0 || k > len(arr) || len(arr) == 0 {
		return nil, false
	}
	var res []int
	h := &MinHeap1{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			res = append(res, heap.Pop(h).(int))
		}
	}
	res = append(res, (*h)...)
	return res, true
}

func main4() {
	arr := []int{6, 5, 3, 2, 8, 10, 9}
	k := 3

	res, _ := FindKSmallest1(arr, k)

	fmt.Println(res)
}
