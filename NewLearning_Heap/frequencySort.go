// package main

// import (
// 	"container/heap"
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 1, 1, 3, 2, 2, 4}
// 	//K := 1

// 	frequencyMap := make(map[int]int)

// 	for i := 0; i < len(arr); i++ {
// 		frequencyMap[arr[i]]++
// 	}

// 	res, _ := FindKSmallest4(frequencyMap)

// 	fmt.Println(res)

// 	for i, num := range res {
// 		if i > 0 {
// 			fmt.Print(",")
// 		}
// 		fmt.Print(num)
// 	}
// 	fmt.Println()

// }

// // MinHeap4 implements a max heap using container/heap
// type MinHeap4 []int

// func (h MinHeap4) Len() int {
// 	return len(h)
// }

// func (h MinHeap4) Less(i, j int) bool {
// 	return h[i] > h[j]
// } // Min heap: smaller elements have priority

// func (h MinHeap4) Swap(i, j int) {
// 	h[i], h[j] = h[j], h[i]
// }

// func (h *MinHeap4) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *MinHeap4) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }

// func FindKSmallest4(frequencyMap map[int]int) ([]int, bool) {
// 	if len(frequencyMap) == 0 {
// 		return nil, false
// 	}
// 	h := &MinHeap4{}
// 	heap.Init(h)
// 	var res []int
// 	//to skip the 0 num
// 	//res = append(res, 0)
// 	for _, value := range frequencyMap {
// 		heap.Push(h, value)
// 	}

// 	for h.Len() > 0 {
// 		m := heap.Pop(h)
// 		res = append(res, m.(int))
// 	}
// 	return res, true
// }

package main

import (
	"container/heap"
	"fmt"
)

// Item represents a number and its frequency
type Item struct {
	num  int // The original number
	freq int // Its frequency
}

// MaxHeap4 implements a max heap for sorting by frequency (descending) and number (ascending)
type MaxHeap4 []Item

func (h MaxHeap4) Len() int { return len(h) }
func (h MaxHeap4) Less(i, j int) bool {
	if h[i].freq == h[j].freq {
		return h[i].num < h[j].num // Sort by number (ascending) if frequencies are equal
	}
	return h[i].freq > h[j].freq // Sort by frequency (descending)
}
func (h MaxHeap4) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap4) Push(x interface{}) { *h = append(*h, x.(Item)) }
func (h *MaxHeap4) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func FindKSmallest4(frequencyMap map[int]int) ([]int, bool) {
	if len(frequencyMap) == 0 {
		return nil, false
	}

	// Initialize max heap
	h := &MaxHeap4{}
	heap.Init(h)

	// Push all number-frequency pairs into the heap
	for num, freq := range frequencyMap {
		heap.Push(h, Item{num: num, freq: freq})
	}

	// Build result array
	var result []int
	for h.Len() > 0 {
		item := heap.Pop(h).(Item)
		// Append the number 'freq' times
		for i := 0; i < item.freq; i++ {
			result = append(result, item.num)
		}
	}

	return result, true
}

func main7() {
	arr := []int{1, 1, 1, 3, 2, 2, 4}

	frequencyMap := make(map[int]int)
	for _, num := range arr {
		frequencyMap[num]++
	}

	res, ok := FindKSmallest4(frequencyMap)
	if !ok {
		fmt.Println("Invalid input")
		return
	}

	// Print in the desired format {1,1,1,2,2,3,4}
	fmt.Print("{")
	for i, num := range res {
		if i > 0 {
			fmt.Print(",")
		}
		fmt.Print(num)
	}
	fmt.Print("}")
	fmt.Println()
}
