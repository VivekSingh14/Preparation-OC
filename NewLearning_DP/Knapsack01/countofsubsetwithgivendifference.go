package main

import "fmt"

//arr is given {1,1,2,3} and given difference 1 now subsets are ->  {1,1,2}-{3} or {1,2}-{1,3} or {1,2}- {1,3} = 3

// s1-s2 = diff
// s1+ s2 = sum
// 2 s1 = diff+sum
// s1 = diff+sum/2
// question is find out the count of subsets having diff+sum/2 only.

// OR

// -------------TARGET SUM-----------target sum----------------

var dp4 [7][11]int

func main() {
	nums := []int{1, 4, 2, 3}
	diff := 1

	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	sum = (diff + sum) / 2

	for i := 0; i < len(nums)+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				dp4[i][j] = 0
			}
			if j == 0 {
				dp4[i][j] = 1
			}
		}
	}

	res := knapsackDP4(nums, sum, len(nums))

	fmt.Println(res)

	//fmt.Println(dp4)

}

func knapsackDP4(nums []int, sum int, n int) int {

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			if nums[i-1] <= j {
				dp4[i][j] = dp4[i-1][j-nums[i-1]] + dp4[i-1][j]
			} else {
				dp4[i][j] = dp4[i-1][j]
			}
		}
	}
	return dp4[n][sum]
}
