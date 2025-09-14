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

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead
	for cur.Next != nil && cur.Next.Next != nil {
		temp := cur.Next
		temp2 := cur.Next.Next.Next

		cur.Next = cur.Next.Next //->2
		cur.Next.Next = temp     //2->1
		temp.Next = temp2        //1->3
		cur = cur.Next.Next
	}
	return dummyHead.Next
}

func createLinkedList(nums []int) *ListNode {
	if len(nums) == 0 {
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
	list := swapPairs(listNode)
	fmt.Println(list)
}
