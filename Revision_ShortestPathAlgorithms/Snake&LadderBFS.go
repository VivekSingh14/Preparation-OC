package main

import "fmt"

func main1() {
	board := [][]int{{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 35, -1, -1, 13, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, 15, -1, -1, -1, -1}}

	n := len(board)

	node := 1

	dir := true

	connections := make([]int, (n*n)+1)

	for idx := range connections {
		connections[idx] = -1
	}

	for i := n - 1; i >= 0; i-- {
		if dir {
			for j := 0; j < n; j++ {
				if board[i][j] != -1 {
					destination := board[i][j]
					connections[node] = destination
				}
				node++
			}
		} else {
			for j := n - 1; j >= 0; j-- {
				if board[i][j] != -1 {
					destination := board[i][j]
					connections[node] = destination
				}
				node++
			}
		}

		dir = !dir
	}

	//Building the graph

	adjlist := make(map[int][]int)

	for i := 1; i <= n*n-1; i++ {
		//all dice numbers
		for count := 1; count <= 6; count++ {
			nbr := count + i

			if nbr <= n*n {
				if connections[nbr] != -1 {
					//has ladder or snake
					adjlist[i] = append(adjlist[i], connections[nbr])
				} else {
					adjlist[i] = append(adjlist[i], nbr)
				}
			}
		}
	}

	// finding the shortest path

	level := 0

	visited := make([]int, (n*n)+1)

	var que []int

	que = append(que, 1)

	visited[1] = 1

	for len(que) != 0 {
		sz := len(que)

		for sz != 0 {
			node := que[0]
			que = que[1:]

			if node == n*n {
				fmt.Println(level)
				return
			}

			for _, nbr := range adjlist[node] {
				if visited[nbr] == 0 {
					visited[nbr] = 1
					que = append(que, nbr)
				}
			}

			sz--
		}
		level++
	}
}
