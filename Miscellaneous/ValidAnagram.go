package main

import "fmt"

func main7() {
	s := "anagram"
	t := "nagaram"

	fmt.Println(isAnagram(s, t))

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}
	var arr [26]int
	var arr2 [26]int

	for i := 0; i < len(s); i++ {
		arr[s[i]-97] = arr[s[i]-97] + 1
		arr2[t[i]-97] = arr2[t[i]-97] + 1
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != arr2[i] {
			return false
		}
	}

	return true
}
