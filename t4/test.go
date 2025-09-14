package main

import (
	"fmt"
	"math"
	"sort"
)

//给定一个整数数组 nums，请你在该数组中任选两个数，保证和最大且为偶数。
//
//示例 1：
//输入：nums = [11,12,13,15,17,2,3,4,7]
//输出：15,17
//
//示例 2：
//输入：nums = [3,2,4,0]
//输出：2,4
//
//示例 3：
//输入：nums = [3,7,9,5]
//输出：7,9

func maxSum(nums []int) []int {
	var num1, num2 []int
	for _, num := range nums {
		if num%2 == 0 {
			num2 = append(num2, num)
		} else {
			num1 = append(num1, num)
		}
	}

	maxSum1, maxSum2 := math.MinInt, math.MinInt
	sort.Ints(num1)
	sort.Ints(num2)
	maxSum1 = num1[len(num1)-1] + num1[len(num1)-2]
	maxSum2 = num2[len(num2)-1] + num2[len(num2)-2]

	if maxSum1 > maxSum2 {
		return []int{num1[len(num1)-1], num1[len(num1)-2]}
	} else {
		return []int{num2[len(num2)-1], num2[len(num2)-2]}
	}
}

func main() {
	nums := []int{11, 12, 13, 15, 17, 2, 3, 4, 7}
	fmt.Println(maxSum(nums))
}
