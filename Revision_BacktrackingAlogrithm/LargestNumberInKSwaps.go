package main

import "fmt"

func main() {
	str := "1234567"
	K := 4

	res := str

	//initially res will be same number so str in last parameter
	findLargestNumberWithK(&str, K, 0, &res)

	fmt.Println(res)

}

func findLargestNumberWithK(str *string, k int, start int, res *string) {
	if k == 0 || start == len(*str)-1 {
		return
	}
	// Find the largest digit to the right
	maxDigit := (*str)[start]

	for i := start + 1; i < len(*str); i++ {
		if (*str)[i] > maxDigit {
			maxDigit = (*str)[i]
		}
	}
	for i := start + 1; i < len(*str); i++ {
		if (*str)[start] < (*str)[i] && (*str)[i] == maxDigit {
			b := []byte(*str)
			b[start], b[i] = b[i], b[start]
			*str = string(b)

			if *str > *res {
				*res = *str
			}

			findLargestNumberWithK(str, k-1, start+1, res)

			b = []byte(*str)
			b[start], b[i] = b[i], b[start]
			*str = string(b)
		}
	}

	//this is for horizontal drifting, if no digit is smaller than  first digit
	findLargestNumberWithK(str, k, start+1, res)

}
