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

//func removeElements(head *ListNode, val int) *ListNode {
//	for head != nil && head.Val == val {
//		head = head.Next
//	}
//	cur := head
//	for cur != nil && cur.Next != nil {
//		if cur.Next.Val == val {
//			cur.Next = cur.Next.Next
//		} else {
//			cur = cur.Next
//		}
//	}
//	return head
//}

func removeElements(head *ListNode, val int) *ListNode {
	dummyHead := &ListNode{Next: head}
	cur := dummyHead

	for cur != nil && cur.Next != nil {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummyHead.Next
}

func main() {
	node := new(ListNode)
	node.Val = 1

	node1 := new(ListNode)
	node1.Val = 7
	node.Next = node1
	node2 := new(ListNode)
	node2.Val = 7
	node1.Next = node2

	node3 := new(ListNode)
	node3.Val = 5
	node2.Next = node3

	node4 := new(ListNode)
	node4.Val = 6
	node3.Next = node4

	removeElements(node, 7)
	//removeNthFromEnd(node, 2)
	fmt.Println(node)
}
