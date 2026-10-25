package main

import (
	"container/heap"
	"fmt"
)

type DSU1 struct {
	parent []int
	rank   []int
}

func NewDSU1(n int) *DSU1 {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &DSU1{
		parent: parent,
		rank:   rank,
	}
}

// Find returns the root/representative of x.
func (d *DSU1) Find1(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find1(d.parent[x])
	}

	return d.parent[x]
}

// Union joins the sets containing a and b.
func (d *DSU1) Union1(a, b int) bool {
	rootA := d.Find1(a)
	rootB := d.Find1(b)

	if rootA == rootB {
		return false
	}

	if d.rank[rootA] < d.rank[rootB] {
		d.parent[rootA] = rootB
		d.rank[rootB] += d.rank[rootA]
	} else {
		d.parent[rootB] = rootA
		d.rank[rootA] += d.rank[rootB]
	}
	return true
}

type Item3 struct {
	node     []int
	distance int
}

type MinHeap3 []Item3

func (h MinHeap3) Len() int {
	return len(h)
}

func (h MinHeap3) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap3) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap3) Push(x interface{}) {
	*h = append(*h, x.(Item3))
}

func (h *MinHeap3) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {

	edges := [][]int{{1, 2, 28}, {1, 6, 10}, {2, 3, 16}, {2, 7, 14}, {3, 4, 12}, {4, 5, 22}, {4, 7, 18}, {5, 6, 25}, {5, 7, 24}}

	n := 7
	dsu := NewDSU1(n + 1)

	h := &MinHeap3{}

	heap.Init(h)

	for idx := range edges {
		heap.Push(h, Item3{node: []int{edges[idx][0], edges[idx][1]}, distance: edges[idx][2]})
		//dsu.Union1(edges[idx][0], edges[idx][1])

	}

	mincost := 0
	for h.Len() > 0 {
		current := heap.Pop(h).(Item3)
		flag := dsu.Union1(current.node[0], current.node[1])
		if flag {
			mincost += current.distance
		}
	}

	fmt.Println(mincost)

}
