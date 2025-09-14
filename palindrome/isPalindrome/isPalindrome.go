package main

import (
	"fmt"
	"strings"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	var s1 []rune
	for _, v := range s {
		if (v >= 'a' && v <= 'z') || (v >= '0' && v <= '9') {
			s1 = append(s1, v)
		}
	}
	left, right := 0, len(s1)-1
	for left <= right {
		if s1[left] != s1[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
