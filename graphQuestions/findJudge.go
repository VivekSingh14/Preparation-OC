package main

import "fmt"

func main10() {
	n := 3

	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}

	fmt.Println(findJudge(n, trust))

}

func findJudge(n int, trusts [][]int) int {

	trustCount := make(map[int]int)

	trustedBy := make(map[int]int)

	for _, trus := range trusts {
		node1, node2 := trus[0], trus[1]
		trustCount[node1]++
		trustedBy[node2]++
	}

	for i := 1; i <= n; i++ {
		if trustedBy[i] == n-1 && trustCount[i] == 0 {
			return i
		}
	}
	return -1
}
