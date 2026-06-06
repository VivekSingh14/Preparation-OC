package main

import (
	"container/heap"
	"fmt"
)

// optimal way find out min cost is always try to sum min elements such as
// given {1,2,3,4,5}
// 1+2 =3    {1+2, 3,4,5}
// 3 + 3    {3+3, 4,5}
// 4+5  {4+5, 6}
// 9+6 {6,9}
// 3+6+9+15 = 33 [min cost]
// which we need to take min heap with top 2 elements
// it will pop top elements and add it into res as well add with each other and push again into min heap
func main1() {
	//arr := []int{1, 2, 3, 4, 5} //o/p: 33
	arr := []int{4, 3, 2, 6} // o/p : 29
	K := 2

	fmt.Println(MinCost(arr, K))
}

// MinHeap implements a max heap using container/heap
type MinHeap6 []int

func (h MinHeap6) Len() int {
	return len(h)
}

func (h MinHeap6) Less(i, j int) bool {
	return h[i] < h[j]
} // Min heap: smaller elements have priority

func (h MinHeap6) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap6) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap6) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func MinCost(arr []int, k int) (int, bool) {
	if k <= 0 || k > len(arr) || len(arr) == 0 {
		return -1, false
	}
	res := 0
	h := &MinHeap6{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
	}

	for h.Len() > 1 {
		sum := 0
		for j := 0; j < k; j++ {
			sum += heap.Pop(h).(int)
		}
		res += sum
		heap.Push(h, sum)

	}
	return res, true
}
