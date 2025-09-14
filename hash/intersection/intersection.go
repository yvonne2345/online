package main

// 两个数组 nums1 和 nums2 ，返回 它们的交集
func intersection(nums1 []int, nums2 []int) []int {
	record := make(map[int]int)
	var result []int
	for _, v := range nums2 {
		record[v] = 1
	}

	for _, v := range nums1 {
		if record[v] == 1 {
			result = append(result, v)
			record[v]++
		}
	}

	return result
}

func main() {
	num1 := []int{4, 9, 5}
	num2 := []int{9, 4, 9, 8, 4}
	//num1 := []int{1, 2, 2, 1}
	//num2 := []int{2, 2}
	//intersection(num1, num2)
	intersect(num1, num2)
	//isHappy(45)
}

func intersect(nums1 []int, nums2 []int) []int {
	//record := make(map[int]int)
	//var result []int
	//if len(nums1) < len(nums2) {
	//	nums1, nums2 = nums2, nums1
	//}
	//for _, v := range nums2 {
	//	record[v]++
	//}
	//
	//for _, v := range nums2 {
	//	if record[v] > 0 {
	//		result = append(result, v)
	//		record[v]--
	//	}
	//}
	//return result

	record := make(map[int]int)
	var result []int
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	for _, v := range nums1 {
		record[v]++
	}

	for _, v := range nums2 {
		if record[v] > 0 {
			result = append(result, v)
			record[v]--
		}
	}
	return result
}

//func isHappy(n int) bool {
//	numMap := make(map[int]bool)
//	for {
//		sum := 0
//		for n != 0 {
//			sum += (n % 10) * (n % 10)
//			n = n / 10
//		}
//		if sum == 1 {
//			return true
//		}
//		if numMap[sum] {
//			return false
//		}
//		numMap[sum] = true
//		n = sum
//	}
//}

func isHappy(n int) bool {
	nmap := make(map[int]bool)
	for {
		sum := 0
		for n != 0 {
			sum += (n % 10) * (n % 10)
			n = n / 10
		}
		if sum == 1 {
			return true
		}
		if nmap[sum] {
			return false
		}
		nmap[sum] = true
		n = sum
	}
}
