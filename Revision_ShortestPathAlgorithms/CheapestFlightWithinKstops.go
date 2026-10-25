package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge1 struct {
	node   int
	weight int
}

type CPGraph struct {
	adjlist map[int][]Edge1
}

func NewCPGraph() *CPGraph {
	return &CPGraph{make(map[int][]Edge1)}
}

func (g *CPGraph) MakeGraph(u int, v int, w int) {
	g.adjlist[u] = append(g.adjlist[u], Edge1{node: v, weight: w})
}

type Item2 struct {
	node     int
	distance int
	stops    int
}

type MinHeap2 []Item2

func (h MinHeap2) Len() int {
	return len(h)
}

func (h MinHeap2) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap2) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap2) Push(x interface{}) {
	*h = append(*h, x.(Item2))
}

func (h *MinHeap2) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main4() {
	n := 4
	flights := [][]int{{0, 1, 100}, {1, 2, 100}, {2, 0, 100}, {1, 3, 600}, {2, 3, 200}}
	src := 0
	dst := 3
	k := 1

	g := NewCPGraph()

	for _, val := range flights {
		g.MakeGraph(val[0], val[1], val[2])
	}

	fmt.Println(cheapestPathWithK(g, src, n+1, dst, k))

}

func cheapestPathWithK(g *CPGraph, src int, n int, dst int, k int) int {
	distance := make([][]int, n)

	for i := range distance {
		distance[i] = make([]int, k+2)
		for j := range distance[i] {
			distance[i][j] = math.MaxInt
		}
	}

	distance[src][0] = 0

	h := &MinHeap2{}

	heap.Init(h)

	heap.Push(h, Item2{node: src, distance: 0, stops: 0})

	for h.Len() > 0 {
		current := heap.Pop(h).(Item2)

		node := current.node
		dis := current.distance
		stop := current.stops

		if node == dst {
			return dis
		}

		if stop == k+1 {
			continue
		}

		for _, nbr := range g.adjlist[node] {
			nbr_node := nbr.node
			nbr_weight := nbr.weight

			newdistance := dis + nbr_weight
			newstops := stop + 1

			if newdistance < distance[nbr_node][newstops] {
				distance[nbr_node][newstops] = newdistance

				heap.Push(h, Item2{
					node:     nbr_node,
					distance: newdistance,
					stops:    newstops,
				})
			}
		}
	}

	return -1
}
