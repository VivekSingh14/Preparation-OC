package main

import (
	"container/heap"
	"fmt"
)

// distance from origin \|---- underoot x^2+y^2
// so instead of using underroot, we will use x^2+y^2 because sequence won't change
// so we can consider key as x^2+y^2
type co struct {
	x, y int
}

func main8() {
	arr := [][]int{{1, 3}, {-2, 2}, {5, 8}, {0, 1}}
	K := 2

	map1 := make(map[int]co)

	for i := 0; i < len(arr); i++ {
		map1[arr[i][0]*arr[i][0]+arr[i][1]*arr[i][1]] = co{arr[i][0], arr[i][1]}
	}
	//fmt.Println(map1)

	res, _ := FindKSmallest5(map1, K)

	for i := 0; i < len(res); i++ {
		temp := map1[res[i]]
		fmt.Println("{", temp.x, ",", temp.y, "}")
	}
}

// MaxHeap4 implements a max heap for sorting by frequency (descending) and number (ascending)
type MaxHeap5 []int

func (h MaxHeap5) Len() int { return len(h) }
func (h MaxHeap5) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h MaxHeap5) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap5) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap5) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKSmallest5(frequencyMap map[int]co, k int) ([]int, bool) {
	if len(frequencyMap) == 0 {
		return nil, false
	}

	// Initialize max heap
	h := &MaxHeap5{}
	heap.Init(h)

	// Push all number-frequency pairs into the heap
	for num := range frequencyMap {
		heap.Push(h, num)
	}

	// Build result array
	//var result []int
	for h.Len() > k {
		heap.Pop(h)
		// Append the number 'freq' times
		// for i := 0; i < item.freq; i++ {
		// 	result = append(result, item.num)
		// }
	}

	return (*h), true
}
