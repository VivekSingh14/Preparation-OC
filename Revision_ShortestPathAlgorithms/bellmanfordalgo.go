package main

import (
	"fmt"
	"math"
)

// type Edge2 struct {
// 	node   int
// 	weight int
// }

// type BFGraph struct {
// 	adjlist map[int][]Edge
// }

// func NewBFGraph() *BFGraph {
// 	return &BFGraph{make(map[int][]Edge)}
// }

// func (g *BFGraph) MakeBFGraph(u int, v int, w int) {
// 	g.adjlist[u] = append(g.adjlist[u], Edge2{node: v, weight: w})
// }

func main5() {
	matrix := [][]int{{1, 2, 4}, {1, 3, 5}, {2, 4, -5}, {3, 4, 3}}
	V := 4

	distance := make([]int, V+1)

	// g := NewBFGraph()

	// for _, val := range list {
	// 	g.MakeBFGraph(val[0], val[1], val[2])
	// }

	for idx := range distance {
		distance[idx] = math.MaxInt
	}

	distance[1] = 0

	for i := 0; i < V-1; i++ {
		for _, edge := range matrix {
			node := edge[0]
			nbr := edge[1]
			wt := edge[2]

			if distance[node] == math.MaxInt {
				continue
			}

			new_dist_of_nbr := distance[node] + wt
			if new_dist_of_nbr < distance[nbr] {
				distance[nbr] = new_dist_of_nbr
			}
		}
	}

	for _, edge := range matrix {
		node := edge[0]
		nbr := edge[1]
		wt := edge[2]

		if distance[node] == math.MaxInt {
			continue
		}

		new_dist_of_nbr := distance[node] + wt
		if new_dist_of_nbr < distance[nbr] {
			distance[nbr] = new_dist_of_nbr
		}
	}

	for i := 0; i < V; i++ {
		if distance[i] == math.MaxInt {
			distance[i] = -200
		}
	}

	fmt.Println(distance)

}
