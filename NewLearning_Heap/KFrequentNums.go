package main

import (
	"container/heap"
	"fmt"
)

func main8() {
	arr := []int{1, 1, 1, 3, 2, 2, 4}
	K := 2

	frequencyMap := make(map[int]int)

	for i := 0; i < len(arr); i++ {
		frequencyMap[arr[i]]++
	}

	res, _ := FindKFrequentNums(frequencyMap, K)

	for _, num := range res {
		for key, value := range frequencyMap {
			if value == num {
				fmt.Print(key, " ")
			}
		}
	}

}

// MinHeap3 implements a max heap using container/heap
type MinHeap3 []int

func (h MinHeap3) Len() int {
	return len(h)
}

func (h MinHeap3) Less(i, j int) bool {
	return h[i] < h[j]
} // Min heap: smaller elements have priority

func (h MinHeap3) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap3) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap3) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKFrequentNums(frequencyMap map[int]int, k int) ([]int, bool) {
	if k <= 0 || k > len(frequencyMap) || len(frequencyMap) == 0 {
		return nil, false
	}
	h := &MinHeap3{}
	heap.Init(h)
	var res []int
	for _, value := range frequencyMap {
		heap.Push(h, value)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	for h.Len() > 0 {
		m := heap.Pop(h)
		res = append(res, m.(int))
	}
	return res, true
}
