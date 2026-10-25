package main

//*************Here i'm traversing the Unweighted UniDirectional Graph*************

import (
	"container/heap"
	"fmt"
)

var arr1 []int

type UnWeightedUniDirectionGraph2 struct {
	adjlist map[int][]int
}

func UWUNIDNewGraph2() *UnWeightedUniDirectionGraph2 {
	return &UnWeightedUniDirectionGraph2{adjlist: make(map[int][]int)}
}

func (g *UnWeightedUniDirectionGraph2) uwuniappend2(u int, v int) {
	arr1 = append(arr1, u)
	// only once because it is uni directional graph
	g.adjlist[u] = append(g.adjlist[u], v)
}

func UWUNItraverse2(g *UnWeightedUniDirectionGraph2, res []int) {

	for val := range res {
		if value, exists := g.adjlist[val]; exists {
			fmt.Println("Vertex: ", val, "is having ", value, "neighbours")
		}
	}
}

func main3() {
	edgelist := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 2}, {1, 3}}

	g2 := UWUNIDNewGraph2()

	for _, v := range edgelist {
		g2.uwuniappend2(v[0], v[1])
	}

	res, _ := FindKSmallest(arr1, 0)

	UWUNItraverse2(g2, res)

}

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
	if len(arr) == 0 {
		return nil, false
	}
	h := &MinHeap{}
	heap.Init(h)

	for _, num := range arr {
		heap.Push(h, num)
	}
	return (*h), true
}
