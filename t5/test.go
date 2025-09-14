package main

import (
	"fmt"
	"sync"
)

//func isValid(s string) {
//	var res []rune
//	for _, v := range s {
//		if v == '(' || v == ')' {
//			res = append(res, v)
//		}
//	}
//	left, right := 0, len(res)-1
//	for left <= right {
//		if s[left] == '(' && s[right] == ')' {
//		}
//	}
//}

func cal(s1 string, s2 string) string {
	str1 := []rune(s1)
	str2 := []rune(s2)

	fmt.Println(str1, str2)
	var a sync.RWMutex
	a.RLock()
	a.Lock()
	return ""
}
