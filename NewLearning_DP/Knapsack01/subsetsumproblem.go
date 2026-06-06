package main

import "fmt"

func main4() {
	fmt.Println("this blocked must be removed")
}

// import "fmt"

// var dp3 [6][12]int

// func main() {
// 	sum := 6
// 	nums := []int{2, 3, 7, 8, 10}

// 	for i := 0; i < len(nums)+1; i++ {
// 		for j := 0; j < sum+1; j++ {
// 			if i == 0 || j == 0 {
// 				dp3[i][j] = 0
// 			}
// 		}
// 	}

// 	res := knapsackDP1(nums, sum, len(nums))

// 	fmt.Println(res)

// 	//fmt.Println(dp3)

// }

// func knapsackDP1(nums []int, sum int, n int) bool {

// 	for i := 1; i < n+1; i++ {
// 		for j := 1; j < sum+1; j++ {
// 			if nums[i-1] <= j {
// 				dp3[i][j] = max((nums[i-1] + dp3[i-1][j-nums[i-1]]), dp3[i-1][j])
// 			} else {
// 				dp3[i][j] = dp3[i-1][j]
// 			}
// 		}
// 	}
// 	return dp3[n][sum] == sum
// }
