package main

import "fmt"

func main8() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	fmt.Println(groupAnagrams(strs))
}

func groupAnagrams(strs []string) [][]string {

	//var res [][]string

	map1 := make(map[int]struct{})

	for i := 0; i < len(strs); i++ {
		var sum int
		for j := 0; j < len(strs[i]); j++ {
			sum += int(strs[i][j])
		}
		map1[sum] = struct{}{}
	}

	fmt.Println(map1)

	resmap := make(map[int][]string)

	for i := 0; i < len(strs); i++ {
		var sum int
		for j := 0; j < len(strs[i]); j++ {
			sum += int(strs[i][j])
		}
		_, ok := map1[sum]
		if ok {
			val := resmap[sum]
			//var arr []string
			val = append(val, strs[i])
			resmap[sum] = val
		}
	}

	return nil

}
