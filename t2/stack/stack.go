package main

import (
	"fmt"
)

// 实现反转输入字符串中（）括号中的字符，例如输入foo(bar)asd，输出foorabasd
// 遍历字符串字符，遇到左括号 ( 时将其推入栈中。
// 遇到右括号 ) 时，从栈中弹出字符直到遇到左括号 (，并将这些字符反转。
// 将反转后的结果插回原位置，继续遍历后续字符。
func reverseInParentheses(input string) string {
	stack := []rune{}
	for _, ch := range input {
		if ch == ')' {
			// 遇到右括号，开始处理反转
			var temp []rune
			// 将括号内的内容出栈
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				temp = append(temp, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			// 弹出左括号 '('
			stack = stack[:len(stack)-1]
			// 将反转后的内容重新入栈
			for _, ch := range temp {
				stack = append(stack, ch)
			}
		} else {
			// 普通字符，直接入栈
			stack = append(stack, ch)
		}
	}
	// 返回拼接成的最终字符串
	return string(stack)
}

func main() {
	// 测试用例
	//input := "foo(bar)asd"
	//result := reverseInParentheses(input)
	//fmt.Println(result) // 输出： "foorabasd"

	arr := []int{1, 2, 3, 1}
	result2 := containsDuplicate(arr)
	fmt.Println(result2) // 输出：true
	//
	//input3 := "aaA"
	//result3 := firstUniqChar(input3)
	//fmt.Println(result3) // 输出：2
}

// 实现判断数组中是否含有重复元素  例如输入1，2，3，1 返回true
// 判断数组中是否有重复元素
func containsDuplicate(nums []int) bool {
	// 创建一个空的 map 用来记录出现过的元素
	seen := make(map[int]bool)

	// 遍历数组
	for _, num := range nums {
		// 如果元素已经在 map 中出现过，则返回 true
		if seen[num] {
			return true
		}
		// 否则将当前元素加入 map
		seen[num] = true
	}

	// 如果遍历完都没有发现重复元素，返回 false
	return false
}

// 在字符串中找到第一个只出现一次的字，并返回它的位置，没有则返回-1，输入aaA，返回2
// 找到第一个只出现一次的字符并返回位置
func firstUniqChar(s string) int {
	// 统计每个字符的出现次数
	count := make(map[rune]int)
	for _, char := range s {
		count[char]++
	}

	// 再次遍历字符串，找到第一个出现一次的字符
	for i, char := range s {
		if count[char] == 1 {
			return i
		}
	}

	// 如果没有找到，返回 -1
	return -1
}
