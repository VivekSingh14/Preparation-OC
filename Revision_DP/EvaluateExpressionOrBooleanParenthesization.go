package main

import "fmt"

func main23() {
	str := "T|T&F^T"

	i := 0

	j := len(str) - 1

	fmt.Println(solveExpress(str, i, j, true))

}

func solveExpress(str string, i int, j int, isTrue bool) int {
	if i > j {
		return 0
	}

	if i == j {
		if isTrue {
			if str[i] == 'T' {
				return 1
			}
			return 0
		}
		if str[i] == 'F' {
			return 1
		}
		return 0
	}

	ans := 0
	for k := i + 1; k <= j-1; k = k + 2 {
		ltrue := solveExpress(str, i, k-1, true)
		lfalse := solveExpress(str, i, k-1, false)
		rtrue := solveExpress(str, k+1, j, true)
		rfalse := solveExpress(str, k+1, j, false)

		if str[k] == '&' {
			if isTrue {
				ans = ans + ltrue*rtrue
			} else {
				ans = ans + (ltrue * rfalse) + (lfalse * rtrue) + (lfalse * rfalse)
			}
		} else if str[k] == '|' {
			if isTrue {
				ans = ans + (ltrue * rtrue) + (ltrue * rfalse) + (lfalse * rtrue)
			} else {
				ans = ans + (lfalse * rfalse)
			}
		} else if str[k] == '^' {
			if isTrue {
				ans = ans + (ltrue * rfalse) + (lfalse * rtrue)
			} else {
				ans = ans + (lfalse * rfalse) + (ltrue * rtrue)
			}
		}

	}

	return ans
}
