package main

import (
	"container/heap"
	"fmt"
)

// MaxHeap implements a max heap using container/heap
type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
} // Max heap: larger elements have priority

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindKthSmallest finds the kth smallest element in the array
func FindKthSmallest(arr []int, k int) (int, bool) {
	if k <= 0 || k > len(arr) || len(arr) == 0 {
		return -1, false
	}

	// Initialize max heap
	h := &MaxHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// The root of the max heap is the kth smallest
	return (*h)[0], true
}

func main() {
	arr := []int{2, 8, 4, 9, 3, 10}

	res, _ := FindKthSmallest(arr, 3)

	fmt.Println(res)

}
