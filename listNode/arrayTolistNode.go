package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转单链表
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

//func main() {
//	listNode := createLinkedList([]int{1, 2, 3, 4, 5})
//	node := createDoubleLinkedList([]int{1, 2, 3, 4, 5})
//	fmt.Println(listNode)
//	fmt.Println(node)
//}

type Node struct {
	Val  int
	Pre  *Node
	Next *Node
}

// 数组转双链表
func createDoubleLinkedList(nums []int) *Node {
	if nums == nil && len(nums) == 0 {
		return &Node{}
	}
	head := &Node{Val: nums[0]}
	cur := head
	for _, num := range nums[1:] {
		temp := &Node{Val: num}
		cur.Next = temp
		temp.Pre = cur
		cur = cur.Next
	}
	return head
}
