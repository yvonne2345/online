package main

import "fmt"

// 删除有序数组中的重复项
//func removeDuplicates(nums []int) int {
//	n:=len(nums)
//	if n<1{
//		return 0
//	}
//	slow := 1
//	for fast := 1; fast < len(nums); fast++ {
//		if nums[fast] != nums[fast-1] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	fmt.Println(nums)
//	return slow
//}

//func moveZeroes(nums []int) {
//	slow := 0
//	for fast := 0; fast < len(nums); fast++ {
//		if nums[fast] != 0 {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	for i := slow; i < len(nums); i++ {
//		nums[i] = 0
//	}
//	fmt.Println(nums)
//}

func moveZeroes(nums []int) {
	slow := 0
	for _, num := range nums {
		if num != 0 {
			nums[slow] = num
			slow++
		}
	}
	for i := slow; i < len(nums); i++ {
		nums[i] = 0
	}
	fmt.Println(nums)
}

//func backspaceCompare(s string, t2 string) bool {
//	resultS := simplify(s)
//	resultT := simplify(t2)
//	return resultS == resultT
//}

func simplify(s string) string {
	if len(s) == 1 {
		return s
	}
	str := []byte(s)
	slow := 0
	for fast := 0; fast < len(str); fast++ {
		if str[fast] != '#' {
			slow++
		} else {
			if slow > 0 {
				slow--
			}
		}
	}
	return string(str[:slow])
}

func backspaceCompare(s string, t string) bool {
	str := simpleString(s)
	tStr := simpleString(t)
	return str == tStr
}

func simpleString(str string) string {
	slow, fast := 0, 0
	s := []byte(str)
	t := make([]byte, 0)
	for fast < len(s) {
		if s[fast] != '#' {
			t = append(t, s[fast])
			slow++
		} else {
			if slow > 0 {
				slow--
				t = t[:slow]
			}
		}
		fast++
	}
	return string(t)
}

//func sortedSquares(nums []int) []int {
//	front := 0
//	backend, count := len(nums)-1, len(nums)-1
//	sortNums := make([]int, count+1)
//	for front <= backend {
//		if nums[front]*nums[front] < nums[backend]*nums[backend] {
//			sortNums[count] = nums[backend] * nums[backend]
//			backend--
//			count--
//		} else {
//			sortNums[count] = nums[front] * nums[front]
//			front++
//			count--
//		}
//	}
//	return sortNums
//}

func main() {
	//nums := []int{0, 0, 1}
	//fmt.Println(removeDuplicates(nums))

	//nums := []int{1, 0, 1, 0, 2}
	//moveZeroes(nums)

	nums := []int{-7, -3, 2, 3, 11}
	sortedSquares(nums)
	//s := "xywrrmp"
	//t2 := "xywrrmu#p"
	//backspaceCompare(s, t2)

	//nums := []int{3, 2, 2, 3}
	//val := 3
	//fmt.Println(removeElement(nums, val))
}

//func removeElement(nums []int, val int) int {
//	slow := 0
//	for fast := 0; fast < len(nums); fast++ {
//		if nums[fast] != val {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums)-1 {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

//func removeDuplicates(nums []int) int {
//	slow := 1
//	for fast := 1; fast < len(nums); fast++ {
//		if nums[fast] != nums[fast-1] {
//			nums[slow] = nums[fast]
//			slow++
//		}
//	}
//	return slow
//}

func removeDuplicates(nums []int) int {
	slow, fast := 1, 1
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func sortedSquares(nums []int) []int {
	slow, fast, count := 0, len(nums)-1, len(nums)-1
	sortNums := make([]int, len(nums))
	for slow <= fast {
		if nums[slow]*nums[slow] < nums[fast]*nums[fast] {
			sortNums[count] = nums[fast] * nums[fast]
			count--
			fast--
		} else {
			sortNums[count] = nums[slow] * nums[slow]
			count--
			slow++
		}
	}
	return sortNums
}
