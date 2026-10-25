package main

import "fmt"

// Unweighted Unidirectional graph
type TopoBFSGraph1 struct {
	adjlist map[int][]int
}

func NewTopoBFSGraph1() *TopoBFSGraph1 {
	return &TopoBFSGraph1{adjlist: make(map[int][]int)}
}

func (g *TopoBFSGraph1) TopoBFSAppend1(u int, v int) {
	g.adjlist[u] = append(g.adjlist[u], v)
}

func TopologicalSortBFS1(s *TopoBFSGraph1, n int) bool {

	// *********** only if number vertex is not given   ***************

	// n := 0
	// for u, neighbors := range s.adjlist {
	// 	if u >= n {
	// 		n = u + 1
	// 	}

	// 	for _, v := range neighbors {
	// 		if v >= n {
	// 			n = v + 1
	// 		}
	// 	}
	// }

	fmt.Println("adjlist: ", s.adjlist)

	indegree := make([]int, n)
	res := make([]int, 0)

	for _, neighbors := range s.adjlist {
		// Make sure vertices with no outgoing edges are also known.
		// if _, ok := indegree[u]; !ok {
		// 	indegree[u] = 0
		// }

		for _, v := range neighbors {
			indegree[v]++
		}
	}

	fmt.Println("indegrees: ", indegree)

	var que []int
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			que = append(que, i)
		}
	}

	for len(que) != 0 {
		node := que[0]
		que = que[1:]

		res = append(res, node)

		for _, nbr := range s.adjlist[node] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				que = append(que, nbr)
			}
		}
	}

	//extra loop to check if cycle exist ?
	for i := 0; i < len(indegree); i++ {
		if indegree[i] != 0 {
			return true
		}
	}

	return false

}

func main18() {

	//cyclic graph
	edgelist := [][]int{{1, 3}, {2, 3}, {4, 3}, {3, 5}, {7, 4}, {5, 7}, {6, 7}}

	//non cyclic graph
	//edgelist := [][]int{{1, 3}, {2, 3}, {3, 4}, {3, 5}, {4, 7}, {5, 7}, {6, 7}}
	g1 := NewTopoBFSGraph1()

	for _, v := range edgelist {
		g1.TopoBFSAppend1(v[0], v[1])
	}

	fmt.Println(TopologicalSortBFS1(g1, 8))

}
