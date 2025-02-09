package main

import "fmt"

func main1() {

	//arr := []int{7, 1, 5, 3, 6, 4}
	//arr := []int{7, 6, 4, 3, 1}
	arr := []int{2, 4, 1}

	fmt.Println(maxProfit(arr))
}

func maxProfit(arr []int) int {
	minprice := arr[0]
	maxprice := 0

	for i := 1; i < len(arr); i++ {
		if minprice > arr[i] {
			minprice = arr[i]
		}
		if arr[i]-minprice > maxprice {
			maxprice = arr[i] - minprice
		}
	}

	return maxprice
}
