package main

import "fmt"

func main() {

	n := 6
	edges := [][]int{{0, 1}, {1, 2}, {2, 3}, {4, 5}}

	fmt.Println(countComponents(n, edges))

}

func countComponents(n int, edges [][]int) int {

	var parent []int
	var rank []int
	for i := 0; i < n; i++ {
		parent = append(parent, i)
		rank = append(rank, 1)
	}

	res := n

	for i := 0; i < len(edges); i++ {
		res -= union(edges[i][0], edges[i][1], parent, rank)
	}

	return res

}

func find(n1 int, parent []int) int {
	res := n1

	for res != parent[res] {
		parent[res] = parent[parent[res]]
		res = parent[res]
	}
	return res
}

func union(n1 int, n2 int, parent []int, rank []int) int {
	p1, p2 := find(n1, parent), find(n2, parent)

	if p1 == p2 {
		return 0
	}

	if rank[p2] > rank[p1] {
		parent[p1] = p2
		rank[p2] += rank[p1]
	} else {
		parent[p2] = p1
		rank[p1] += rank[p2]
	}
	return 1
}
