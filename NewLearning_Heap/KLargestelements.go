package main

import (
	"container/heap"
	"fmt"
)

// MinHeap implements a max heap using container/heap
type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
} // Min heap: smaller elements have priority

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKSmallest(arr []int, k int) ([]int, bool) {
	if k <= 0 || k > len(arr) || len(arr) == 0 {
		return nil, false
	}
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	return (*h), true
}

func main2() {
	arr := []int{2, 8, 4, 9, 3, 10}

	res, _ := FindKSmallest(arr, 3)

	fmt.Println(res)

}
