package main

import (
	"sort"
)

// 数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
//func twoSum(nums []int, target int) []int {
//	numMap := make(map[int]int)
//	for index, num := range nums {
//		value := target - num
//		valueIndex, ok := numMap[value]
//		if ok {
//			return []int{index, valueIndex}
//		} else {
//			numMap[num] = index
//		}
//	}
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	nmap := make(map[int]int)
	for k, v := range nums {
		value := target - v
		valueIndex, ok := nmap[value]
		if ok {
			return []int{k, valueIndex}
		} else {
			nmap[v] = k
		}
	}
	return []int{}
}

func main() {
	//num := []int{11, 7, 2, 15}
	//target := 9
	//twoSum(num, target)
	//nums := []int{-1, 0, 1, 2, -1, -4}
	//threeSum(nums)
	//nums := []int{1, 0, -1, 0, -2, 2}
	nums := []int{-2, -1, -1, 1, 1, 2, 2}
	target := 0
	fourSum(nums, target)
	//nums1 := []int{1, 2}
	//nums2 := []int{-2, -1}
	//nums3 := []int{-1, 2}
	//nums4 := []int{0, 2}
	//fourSumCount(nums1, nums2, nums3, nums4)
}

// 三数之和
//func threeSum(nums []int) [][]int {
//	sort.Ints(nums)
//	result := [][]int{}
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			return result
//		}
//		//去重，当某一位数是i时，第一次遍历已经将所有有关于i的结果输出，则下一位是即使也是i也不影响结果，所以只需要用countine跳出循环就可以了。
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		left := i + 1
//		right := len(nums) - 1
//		for right > left {
//			if nums[i]+nums[left]+nums[right] > 0 {
//				right--
//			} else if nums[i]+nums[left]+nums[right] < 0 {
//				left++
//			} else {
//				//先存结果，再去重
//				result = append(result, []int{nums[i], nums[left], nums[right]})
//				//n2n3是一组解中的左值和右值，这里去重是通过去掉源数组中和左值右值相等的数来实现
//				//for l < r && nums[l] == n2 {
//				//	l++
//				//}
//				//for l < r && nums[r] == n3 {
//				//	r--
//				//}
//
//				for right > left && nums[right] == nums[right-1] {
//					right--
//				}
//				for right > left && nums[left] == nums[left+1] {
//					left++
//				}
//				right--
//				left++
//			}
//		}
//	}
//	return result
//}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := range nums {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			if nums[i]+nums[left]+nums[right] > 0 {
				right--
			} else if nums[i]+nums[left]+nums[right] < 0 {
				left++
			} else {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for right > left && nums[right] == nums[right-1] {
					right--
				}
				for right > left && nums[left] == nums[left+1] {
					left++
				}
				right--
				left++
			}
		}
	}
	println(ans)
	return ans
}

// 四数之和
//func fourSum(nums []int, target int) [][]int {
//	sort.Ints(nums)
//	result := [][]int{}
//	for i := 0; i < len(nums)-3; i++ {
//		if i > 0 && nums[i] == nums[i-1] {
//			continue
//		}
//		for j := i + 1; j < len(nums)-2; j++ {
//			if j > i+1 && nums[j] == nums[j-1] {
//				continue
//			}
//			left := j + 1
//			right := len(nums) - 1
//			for right > left {
//				if nums[i]+nums[j]+nums[left]+nums[right] > target {
//					right--
//				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
//					left++
//				} else {
//					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})
//					for right > left && nums[right] == nums[right-1] {
//						right--
//					}
//					for right > left && nums[left] == nums[left+1] {
//						left++
//					}
//					right--
//					left++
//				}
//			}
//		}
//
//	}
//	return result
//}

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		//if nums[i] > target {
		//	return ans
		//}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			left := j + 2
			right := len(nums) - 1
			for left < right {
				if nums[i]+nums[j]+nums[left]+nums[right] > target {
					right--
				} else if nums[i]+nums[j]+nums[left]+nums[right] < target {
					left++
				} else {
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					for right > left && nums[left] == nums[left+1] {
						left++
					}
					for right > left && nums[right] == nums[right-1] {
						right--
					}
					right--
					left++
				}
			}
		}
	}
	return ans
}

// 四数相加2-数组
//
//	func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
//		m := make(map[int]int)
//		count := 0
//		for _, a := range nums1 {
//			for _, b := range nums2 {
//				//记录该值次数
//				m[a+b]++
//			}
//		}
//		for _, c := range nums3 {
//			for _, d := range nums4 {
//				value, ok := m[-c-d]
//				if ok {
//					count += value
//				}
//			}
//		}
//		return count
//	}
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	nmap := make(map[int]int)
	for _, a := range nums1 {
		for _, b := range nums2 {
			nmap[a+b]++
		}
	}
	count := 0
	for _, c := range nums3 {
		for _, d := range nums4 {
			num, ok := nmap[-c-d]
			if ok {
				count += num
			}
		}
	}
	return count
}
