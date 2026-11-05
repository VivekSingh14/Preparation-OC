package main

import (
	"fmt"
	"strings"
)

func main2() {
	s := "A man, a plan, a canal: Panama"

	s = strings.ToLower(s)
	newstr := ""

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9') {
			newstr += string(s[i])
		}
	}

	fmt.Println(isPalindrome(newstr))

}

func isPalindrome(str string) bool {
	first := 0

	last := len(str) - 1

	for first < last {
		if str[first] != str[last] {
			return false
		}
		first++
		last--

	}
	return true
}
