package detectCycle

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

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil { //设定快慢指针，同时快指针比满指针快一步
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow { //快慢指针相等，找到环入口
			temp1 := fast
			for temp1 != head {
				temp1 = temp1.Next
				head = head.Next
			}
			return temp1
		}
	}
	return nil
}
