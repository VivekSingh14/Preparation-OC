package main

import "fmt"

func main12() {

	str1 := "xyabcde"

	str2 := "lmabcfce"

	m := len(str1)

	n := len(str2)

	dp3 := make([][]int, 1001)

	for x := range dp3 {
		dp3[x] = make([]int, 1001)
	}

	maxlen := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp3[i][j] = 1 + dp3[i-1][j-1]
				if dp3[i][j] > maxlen {
					maxlen = dp3[i][j]
				}
			} else {
				dp3[i][j] = 0
			}
		}
	}

	fmt.Println(maxlen)
}
