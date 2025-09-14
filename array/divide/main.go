package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left++
		} else if target < nums[mid] {
			right--
		} else {
			return mid
		}
	}
	return -1
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target > nums[mid] {
			left++
		} else if target < nums[mid] {
			right--
		} else {
			return mid
		}
	}
	return right + 1
}

func mySqrt(x int) int {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid > x {
			right--
		} else if mid*mid < x {
			left++
		} else {
			return mid
		}
	}
	return left - 1
}

func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (left + right) / 2
		if mid*mid < num {
			left = mid + 1
		} else if mid*mid > num {
			right = mid - 1
		} else {
			return true
		}
	}
	return false
}

func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	//target := 9
	//nums := []int{9}
	//target := 9
	//fmt.Println(search(nums, target))

	//nums := []int{1, 3, 5, 6}
	//target := 4
	//fmt.Println(searchInsert(nums, target))

	//x := 2147395600
	//fmt.Println(mySqrt(x))

	//num := 16
	//fmt.Println(isPerfectSquare(num))

	nums := []int{5, 7, 7, 8, 8, 10}
	target := 6
	fmt.Println(searchRange(nums, target))
}

func searchRange(nums []int, target int) []int {
	left := getLeft(nums, target)
	right := getRight(nums, target)
	if left == -2 || right == -2 {
		return []int{-1, -1}
	}
	if right-left > 1 {
		return []int{left + 1, right - 1}
	}
	return []int{-1, -1}
}

func getRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1

		} else {
			left = mid + 1
			border = left
		}
	}
	return border
}

func getLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	border := -2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
			border = right
		} else {
			left = mid + 1
		}
	}
	return border
}
