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

func reverseList(head *ListNode) *ListNode {
	cur := head
	//pre := &ListNode{}
	var pre *ListNode

	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre

}

func createLinkedList(nums []int) *ListNode {
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

func main() {
	listNode := createLinkedList([]int{1, 2, 3, 4, 5})
	list := reverseList(listNode)
	fmt.Println(list)
}
