package main

import "fmt"

//type MyLinkedList struct {
//	Size      int
//	dummyHead *ListNode
//}
//
//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
//
//func Constructor() MyLinkedList {
//	newNode := &ListNode{-1, nil}
//	//newNode := &ListNode{}
//	return MyLinkedList{
//		Size:      0,
//		dummyHead: newNode,
//	}
//}
//
//func (this *MyLinkedList) Get(index int) int {
//	if index < 0 || index >= this.Size {
//		return -1
//	}
//	cur := this.dummyHead
//	for i := 0; i <= index; i++ {
//		cur = cur.Next
//	}
//	return cur.Val
//}

//func (this *MyLinkedList) AddAtHead(val int) {
//	temp := &ListNode{Val: val}
//	cur := this.dummyHead
//	temp.Next = cur.Next
//	this.dummyHead.Next = temp
//	this.Size++
//}

func (this *MyLinkedList) AddAtHead(num int) {
	temp := &ListNode{Val: num}
	cur := this.dummyHead
	temp.Next = cur.Next
	cur.Next = temp
	this.Size++
}

//func (this *MyLinkedList) AddAtTail(val int) {
//	cur := this.dummyHead
//	temp := &ListNode{Val: val}
//	for cur.Next != nil {
//		cur = cur.Next
//	}
//	cur.Next = temp
//	this.Size++
//}

func (this *MyLinkedList) AddAtTail(val int) {
	temp := &ListNode{Val: val}
	cur := this.dummyHead
	for cur != nil && cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = temp
	this.Size++
}

//
//func (this *MyLinkedList) AddAtIndex(index int, val int) {
//	if index > this.Size {
//		return
//	} else if index < 0 {
//		index = 0
//	}
//
//	temp := &ListNode{Val: val}
//	cur := this.dummyHead
//	for i := 0; i < index; i++ {
//		cur = cur.Next
//	}
//	temp.Next = cur.Next
//	cur.Next = temp
//	this.Size++
//}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	temp := &ListNode{Val: val}
	cur := this.dummyHead
	if index < 0 {
		index = 0
	} else if index > this.Size {
		return
	} //注意大于size判断
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	temp.Next = cur.Next
	cur.Next = temp
	this.Size++
}

//func (this *MyLinkedList) DeleteAtIndex(index int) {
//	if index > this.Size {
//		return
//	} else if index < 0 {
//		index = 0
//	}
//	cur := this.dummyHead
//	for i := 0; i < index; i++ {
//		cur = cur.Next
//	}
//	cur.Next = cur.Next.Next
//	this.Size--
//}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.dummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next //注意当前指针位置
	this.Size--
}

func main() {
	list := Constructor() // 初始化链表
	list.AddAtHead(1)     // 在头部添加元素
	list.AddAtHead(3)     // 在头部添加元素
	list.AddAtIndex(1, 0) // 在指定位置添加元素
	list.Get(1)
	list.DeleteAtIndex(1)
	list.Get(1)
	list.AddAtTail(777) // 在尾部添加元素
	list.AddAtTail(242) // 在尾部添加元素
	fmt.Println(list)
}

func (this *MyLinkedList) Get(i int) int {
	if i >= this.Size || i < 0 {
		return 0
	}
	cur := this.dummyHead
	for index := 0; index <= i; index++ {
		cur = cur.Next
	}
	return cur.Val
}

func Constructor() MyLinkedList {
	node := &ListNode{}
	return MyLinkedList{
		Size:      0,
		dummyHead: node,
	}
}

type MyLinkedList struct {
	Size      int
	dummyHead *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}
