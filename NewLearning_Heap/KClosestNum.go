package main

import (
	"container/heap"
	"fmt"
	"math"
)

type mapping struct {
	closest, num int
}

// MaxHeap2 implements a max heap for mapping structs
type MaxHeap2 []mapping

func (h MaxHeap2) Len() int {
	return len(h)
}

func (h MaxHeap2) Less(i, j int) bool {
	return h[i].closest > h[j].closest // Max heap: larger differences have priority
}

func (h MaxHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap2) Push(x interface{}) {
	*h = append(*h, x.(mapping))
}

func (h *MaxHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// FindKthSmallest2 finds the k closest numbers to X
func FindKthSmallest2(arr []mapping, k int) ([]mapping, bool) {
	if k <= 0 || k > len(arr) || len(arr) == 0 {
		return nil, false
	}

	// Initialize max heap
	h := &MaxHeap2{}
	heap.Init(h)

	for _, m := range arr {
		heap.Push(h, m)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Collect the k closest numbers
	// result := make([]int, 0, k)
	// for h.Len() > 0 {
	// 	m := heap.Pop(h).(mapping)
	// 	result = append(result, m.num)
	// }

	//or

	return (*h), true // or return result, true
}

// find the k closest numbers to X from given array
// so in this case 6,7,8 will returned cause 7-6 = 1, 7-7 = 0, |7-8| = 1
func main6() {
	arr := []int{5, 6, 7, 8, 9}
	K := 3
	X := 7
	var temp []mapping
	for i := 0; i < len(arr); i++ {
		tempDiff := int(math.Abs(float64(X - arr[i])))
		temp = append(temp, mapping{tempDiff, arr[i]})
	}

	result, ok := FindKthSmallest2(temp, K)
	if ok {
		fmt.Println(result) // Output: [6 7 8] or [{1 6} {1 8} {0 7}]
	} else {
		fmt.Println("Invalid input")
	}
}
