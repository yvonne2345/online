package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	curA, curB := headA, headB
	for curA != curB {
		if curA != nil {
			curA = curA.Next
		} else {
			curA = headB
		}

		if curB != nil {
			curB = curB.Next
		} else {
			curB = headA
		}
	}
	return curA
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	curA, curB := headA, headB
//	lenA, lenB := 0, 0
//	for curA != nil {
//		lenA++
//		curA = curA.Next
//	}
//	for curB != nil {
//		lenB++
//		curB = curB.Next
//	}
//	pA, pB := headA, headB
//	if lenA > lenB {
//		for i := 0; i < lenA-lenB; i++ {
//			pA = pA.Next
//		}
//	} else {
//		for i := 0; i < lenB-lenA; i++ {
//			pB = pB.Next
//		}
//	}
//	for pA != pB {
//		pA = pA.Next
//		pB = pB.Next
//	}
//	return pB
//}

func main() {
	listNode1 := createListNode([]int{4, 1, 8, 4, 5})
	listNode2 := createListNode([]int{5, 0, 1, 8, 4, 5})
	node := getIntersectionNode(listNode1, listNode2)
	fmt.Println(node)
}

func createListNode(nums []int) *ListNode {
	if nums == nil && len(nums) == 0 {
		return &ListNode{}
	}
	head := &ListNode{Val: nums[0]}
	cur := head
	for i, num := range nums[1:] {
		if i < len(nums) {
			cur.Next = &ListNode{Val: num}
			cur = cur.Next
		}
	}
	return head
}
