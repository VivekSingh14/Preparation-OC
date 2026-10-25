package main

import "fmt"

// Unweighted Unidirectional graph
type LCBFSGraph struct {
	adjlist map[int][]int
}

func NewLCBFSGraph() *LCBFSGraph {
	return &LCBFSGraph{adjlist: make(map[int][]int)}
}

func (g *LCBFSGraph) LCBFSAppend(u int, v int) {

	g.adjlist[u] = append(g.adjlist[u], v)
}

func main() {

	colors := "abaca"
	edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {3, 4}}

	g1 := NewLCBFSGraph()

	for _, j := range edges {
		g1.LCBFSAppend(j[0], j[1])
	}

	fmt.Println(largestColorPath(colors, g1))

}

func largestColorPath(colors string, g1 *LCBFSGraph) int {

	ans := 0
	n := len(colors)
	indegree := make([]int, n)
	cnt := make([][]int, n)

	for val := range cnt {
		cnt[val] = make([]int, 26)
	}

	for _, neighbors := range g1.adjlist {
		for _, v := range neighbors {
			indegree[v]++
		}
	}

	var que []int
	for i := 0; i < len(indegree); i++ {
		if indegree[i] == 0 {
			que = append(que, i)
		}
	}

	processed := 0

	for len(que) != 0 {
		node := que[0]
		que = que[1:]

		processed += 1

		//Very important to convert char to int
		test := colors[node] - 'a'
		cnt[node][int(test)] += 1

		ans = max(ans, cnt[node][int(test)])

		for _, nbr := range g1.adjlist[node] {
			indegree[nbr]--
			if indegree[nbr] == 0 {
				que = append(que, nbr)
			}

			for j := 0; j < 26; j++ {
				cnt[nbr][j] = max(cnt[nbr][j], cnt[node][j])
			}
		}

	}

	if processed == n {
		return ans
	}
	return -1
}
