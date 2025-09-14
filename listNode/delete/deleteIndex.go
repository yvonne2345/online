package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
//func removeNthFromEnd(head *ListNode, n int) *ListNode {
//	dummyHead := &ListNode{Next: head}
//	slow, fast := dummyHead, dummyHead
//	for i := 0; i <= n; i++ {
//		fast = fast.Next //fast移动n
//	}
//	for fast.Next != nil {
//		//fast、slow同步移动直到fast为nil
//		fast = fast.Next
//		slow = slow.Next
//	}
//	slow.Next = slow.Next.Next
//	return dummyHead.Next
//}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next
	return dummyHead.Next
}
