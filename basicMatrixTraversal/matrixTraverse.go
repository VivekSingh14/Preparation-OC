package main

import "fmt"

func main1() {

	arr := make([][]int, 4)

	fmt.Println("before: length of arr's first row ", len(arr[0]))

	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, 5)
	}

	fmt.Println("after: length of arr's first row ", len(arr[0]))

	fmt.Println("first column")
	for rows := range arr {
		arr[rows][0] = rows + 1

		// arr[rows][0] = rows
		// fmt.Println(arr)

	}

	for rows := range arr {
		fmt.Println(arr[rows])
	}

	fmt.Println("last column")
	for rows := range arr {
		arr[rows][len(arr[rows])-1] = len(arr[rows]) - 1

	}

	for rows := range arr {
		fmt.Println(arr[rows])
	}

	fmt.Println("first row")

	for col := range arr[0] {
		arr[0][col] = col
	}

	for rows := range arr {
		fmt.Println(arr[rows])
	}

	fmt.Println("last row")
	for col := range arr[len(arr)-1] {
		arr[len(arr)-1][col] = col
	}

	for rows := range arr {
		fmt.Println(arr[rows])
	}

	//fmt.Println(arr)
}
