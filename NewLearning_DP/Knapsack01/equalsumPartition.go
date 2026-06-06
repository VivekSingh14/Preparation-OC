package main

// import "fmt"

// var dp3 [6][12]int

// // this problem check if equal subset of sum are given in num array or not.  such as  1,5,11,5 ->  1,5,5  == 11
// func main() {
// 	nums := []int{1, 5, 11, 5}
// 	sum := 0
// 	for i := 0; i < len(nums); i++ {
// 		sum += nums[i]
// 	}

// 	if sum%2 != 0 {
// 		fmt.Println("Not present")
// 		return
// 	}

// 	sum = sum / 2

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
