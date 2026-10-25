package main

import (
	"container/heap"
	"fmt"
	"math"
)

var neighbours = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

type Item1 struct {
	node     []int
	distance int
}

type MinHeap1 []Item1

func (h MinHeap1) Len() int {
	return len(h)
}

func (h MinHeap1) Less(i, j int) bool {
	return h[i].distance < h[j].distance
}

func (h MinHeap1) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap1) Push(x interface{}) {
	*h = append(*h, x.(Item1))
}

func (h *MinHeap1) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main3() {
	heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}

	n := len(heights)
	m := len(heights[0])

	distance := make([][]int, n)

	for idx := range distance {
		distance[idx] = make([]int, m)
		for idy := range distance[idx] {
			distance[idx][idy] = math.MaxInt
		}
	}

	distance[0][0] = 0

	h := &MinHeap1{}

	heap.Init(h)

	heap.Push(h, Item1{node: []int{0, 0}, distance: 0})

	for h.Len() > 0 {
		current := heap.Pop(h).(Item1)
		temp_node := current.node
		dis := current.distance
		x := temp_node[0]
		y := temp_node[1]

		for k := 0; k < 4; k++ {
			ii := x + neighbours[k][0]
			jj := y + neighbours[k][1]

			if ii < 0 || jj < 0 || ii >= n || jj >= m {
				continue
			}

			new_dis_of_nbr := max(dis, int(math.Abs(float64(heights[x][y])-float64(heights[ii][jj]))))

			if new_dis_of_nbr < distance[ii][jj] {
				distance[ii][jj] = new_dis_of_nbr
				heap.Push(h, Item1{node: []int{ii, jj}, distance: distance[ii][jj]})
			}
		}
	}

	fmt.Println(distance[n-1][m-1])

}
