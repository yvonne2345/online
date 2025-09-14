package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//func reverseString(s []byte) {
//	if s == nil {
//		return
//	}
//	slow, fast := 0, len(s)-1
//	for slow < fast {
//		temp := s[slow]
//		s[slow] = s[fast]
//		s[fast] = temp
//		//s[slow],s[fast]=s[fast],s[slow]
//		slow++
//		fast--
//	}
//}

func reverseString(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

func reverseStr(s string, k int) string {
	if len(s) == 0 {
		return ""
	}
	str := []byte(s)
	for i := 0; i < len(str); i = i + 2*k {
		if i+k < len(s) {
			reverse(str[i : i+k])
		} else {
			reverse(str[i:])
		}
	}
	return string(str)
}

func reverse(s []byte) {
	slow, fast := 0, len(s)-1
	for slow < fast {
		s[slow], s[fast] = s[fast], s[slow]
		slow++
		fast--
	}
}

func main() {
	//str := "abcd"
	//reverseStr(str, 4)

	//给定一个字符串 s，它包含小写字母和数字字符，请编写一个函数，将字符串中的字母字符保持不变，而将每个数字字符替换为number。 例如，对于输入字符串 "a1b2c3"，函数应该将其转换为 "anumberbnumbercnumber"。
	//str := "a1b2c3"
	//ss := []byte(str)//给定

	//从键盘输入
	//var ss []byte
	//fmt.Scanln(&ss) //fmt.Scan(&ss)
	//result := make([]string, 0)
	//for _, value := range ss {
	//	if value >= '0' && value <= '9' {
	//		result = append(result, "number")
	//	} else {
	//		result = append(result, string(value))
	//	}
	//}
	//
	//fmt.Println(strings.Join(result, ""))

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ss := scanner.Bytes()
		result := make([]string, 0)
		for _, value := range ss {
			if value >= '0' && value <= '9' {
				result = append(result, "number")
			} else {
				result = append(result, string(value))
			}
		}
		fmt.Println(strings.Join(result, ""))
	}
}
