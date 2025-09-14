package main

import "math"

//func totalFruit(fruits []int) int {
//	left, right, res := 0, 0, 0
//	n := len(fruits)
//	m := make(map[int]int)
//	for right < n {
//		m[fruits[right]]++
//		if len(m) > 2 {
//			m[fruits[left]]--
//			if m[fruits[left]] == 0 {
//				delete(m, fruits[left])
//			}
//			left++
//		}
//		res = max(res, right-left+1)
//		right++
//	}
//	return res
//}
//func max(res int, i int) int {
//	if res > i {
//		return res
//	}
//	return i
//}

func minWindow(s string, t string) string {
	left, right, match := 0, 0, 0
	need := make(map[byte]int)
	win := make(map[byte]int)
	minLen := math.MaxInt
	start := 0
	for i := range t {
		need[t[i]]++
	}

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			win[c]++
			if win[c] == need[c] {
				match++
			}
		}

		for match == len(need) {
			if right-left < minLen {
				minLen = right - left
				start = left
			}
			c = s[left]
			left++
			if _, ok := need[c]; ok {
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}
	if minLen == math.MaxInt {
		return ""
	}
	return s[start : start+minLen]
}

//func minSubArrayLen(target int, nums []int) int {
//	fast, slow := 0, 0
//	sum := 0
//	minLen := math.MaxInt
//	for fast < len(nums) {
//		sum += nums[fast]
//		for sum >= target {
//			//对应外层判断用if,当前for可以省略以下代码
//			//for slow <= fast && sum >= target {
//			//	sum -= nums[slow]
//			//	minLen = min(minLen, fast-slow+1)
//			//	slow++
//			//}
//			sum -= nums[slow]
//			minLen = min(minLen, fast-slow+1)
//			slow++
//		}
//		fast++
//	}
//
//	if minLen == math.MaxInt {
//		return 0
//	}
//	return minLen
//}

func min(res int, i int) int {
	if res > i {
		return i
	}
	return res
}

func minSubArrayLen(target int, nums []int) int {
	slow, fast := 0, 0
	minLen := math.MaxInt
	sum := 0
	for fast < len(nums) {
		sum += nums[fast]
		for sum >= target {
			sum = sum - nums[slow]
			minLen = min(minLen, fast-slow+1)
			slow++
		}
		fast++
	}
	if minLen == math.MaxInt {
		minLen = 0
	}
	return minLen
}

func main() {
	//nums := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
	//totalFruit(nums)

	nums := []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8}
	target := 15
	//nums := []int{1, 4, 4}
	//target := 4
	minSubArrayLen(target, nums)
}
