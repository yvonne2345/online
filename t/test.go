package main

// 在排序数组中查找元素的第一个和最后一个位置
func main() {
	//nums := []int{-1, 0, 3, 5, 9, 12}
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	searchRange(nums, target)
}

//func searchRange(nums []int, target int) []int {
//	leftBorder := getLeftBorder(nums, target)
//	rightBorder := getRightBorder(nums, target)
//	if leftBorder == -2 && rightBorder == -2 {
//		return []int{-1, -1}
//	}
//	if (rightBorder - leftBorder) > 1 {
//		return []int{leftBorder + 1, rightBorder - 1}
//	}
//	return []int{-1, -1}
//
//}
//
//func getLeftBorder(nums []int, target int) (leftBorder int) {
//	//左
//	left := 0
//	right := len(nums) - 1
//	leftBorder = -2
//	for left <= right {
//		mid := (left + right) / 2
//		if nums[mid] >= target {
//			right = mid - 1
//			leftBorder = right
//		} else {
//			left = mid + 1
//		}
//	}
//	return leftBorder
//}
//
//func getRightBorder(nums []int, target int) (rightBorder int) {
//	left := 0
//	right := len(nums) - 1
//	rightBorder = -2
//	for left <= right {
//		mid := (left + right) / 2
//		if nums[mid] > target {
//			right = mid - 1
//		} else {
//			left = mid + 1
//			rightBorder = left
//		}
//	}
//	return rightBorder
//}

func searchRange(nums []int, target int) []int {
	leftBorder := leftBorder(nums, target)
	rightBorder := rightBorder(nums, target)
	if leftBorder == -2 && rightBorder == -2 {
		return []int{-1, -1}
	}
	if (rightBorder - leftBorder) > 1 {
		return []int{leftBorder + 1, rightBorder - 1}
	}
	return []int{-1, -1}
}

func leftBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	leftBorder := -2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
			leftBorder = right
		}
	}
	return leftBorder
}

func rightBorder(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	rightBorder := -2
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
			rightBorder = left
		} else {
			right = mid - 1
		}
	}
	return rightBorder
}
