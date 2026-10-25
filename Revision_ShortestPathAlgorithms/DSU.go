package main

import "fmt"

type DSU struct {
	parent []int
	rank   []int
}

func NewDSU(n int) *DSU {
	parent := make([]int, n)
	rank := make([]int, n)

	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 1
	}

	return &DSU{
		parent: parent,
		rank:   rank,
	}
}

// Find returns the root/representative of x.
func (d *DSU) Find(x int) int {
	if d.parent[x] != x {
		d.parent[x] = d.Find(d.parent[x])
	}

	return d.parent[x]
}

// Union joins the sets containing a and b.
func (d *DSU) Union(a, b int) {
	rootA := d.Find(a)
	rootB := d.Find(b)

	if rootA == rootB {
		return
	}

	if d.rank[rootA] < d.rank[rootB] {
		d.parent[rootA] = rootB
		d.rank[rootB] += d.rank[rootA]
	} else {
		d.parent[rootB] = rootA
		d.rank[rootA] += d.rank[rootB]
	}
}

func main7() {
	dsu := NewDSU(5)

	dsu.Union(0, 1)
	dsu.Union(1, 2)
	dsu.Union(3, 4)

	fmt.Println(dsu.Find(0) == dsu.Find(2)) // true
	fmt.Println(dsu.Find(0) == dsu.Find(3)) // false

	dsu.Union(2, 3)

	fmt.Println(dsu.Find(0) == dsu.Find(4)) // true
}
