package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	node   int
	weight int
}

type DKGraph struct {
	adjlist map[int][]Edge
}

func NewDKGraph() *DKGraph {
	return &DKGraph{make(map[int][]Edge)}
}

func (g *DKGraph) MakeGraph(u int, v int, w int) {
	g.adjlist[u] = append(g.adjlist[u], Edge{node: v, weight: w})
}

type Item struct {
	node     int
	distance int
}

type MinHeap []Item

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main2() {
	// list := [][]int{{0, 1, 1}, {0, 2, 6}, {1, 2, 3}, {1, 0, 1}, {2, 1, 3}, {2, 0, 6}}
	// src := 2
	// n := 3

	list := [][]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}
	n := 4
	src := 2

	g := NewDKGraph()

	for _, val := range list {
		g.MakeGraph(val[0], val[1], val[2])
	}

	fmt.Println(dijkstra(g, src, n+1))

}

func dijkstra(g *DKGraph, src int, n int) []int {
	distance := make([]int, n)

	for i := range distance {
		distance[i] = math.MaxInt
	}

	distance[src] = 0

	h := &MinHeap{}

	heap.Init(h)

	heap.Push(h, Item{node: src, distance: 0})

	for h.Len() > 0 {
		current := heap.Pop(h).(Item)

		node := current.node
		dis := current.distance

		if dis > distance[node] {
			continue
		}

		for _, nbr := range g.adjlist[node] {
			nbr_node := nbr.node
			nbr_weight := nbr.weight

			newdistance := dis + nbr_weight

			if newdistance < distance[nbr_node] {
				distance[nbr_node] = newdistance

				heap.Push(h, Item{
					node:     nbr_node,
					distance: newdistance,
				})
			}
		}
	}
	return distance
}
