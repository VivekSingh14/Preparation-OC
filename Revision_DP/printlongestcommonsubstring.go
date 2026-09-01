package main

import "fmt"

func main14() {

	str1 := "abde"

	str2 := "abfce"

	m := len(str1)

	n := len(str2)

	dp3 := make([][]int, 11)

	for x := range dp3 {
		dp3[x] = make([]int, 11)
	}

	maxlen := 0

	x := 0
	y := 0

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp3[i][j] = 1 + dp3[i-1][j-1]
				if dp3[i][j] > maxlen {
					maxlen = dp3[i][j]
					x = i
					y = j
				}
			} else {
				dp3[i][j] = 0
			}
		}
	}

	var res string
	limit := x - maxlen

	for x > limit && y > limit {
		if str1[x-1] == str2[y-1] {
			res += string(str1[x-1])
			x--
			y--
		}
	}

	res = Reverse1(res)

	fmt.Println(res)

}

func Reverse1(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
