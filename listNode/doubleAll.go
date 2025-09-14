package main

//type Node struct {
//	Val  int
//	Pre  *Node
//	Next *Node
//}

type MyLinkedList struct {
	Size      int
	DummyHead *Node
	DummyTail *Node
}

func Constructor() MyLinkedList {
	dummyHead := &Node{-1, nil, nil}
	dummyTail := &Node{-2, nil, dummyHead}
	dummyHead.Next = dummyTail //指向尾结点

	return MyLinkedList{0, dummyHead, dummyTail}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index > this.Size {
		return -1
	}
	cur := this.DummyHead
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	cur := this.DummyHead
	temp := &Node{Val: val}
	temp.Next = cur.Next
	temp.Pre = cur
	cur.Next.Pre = temp
	cur.Next = temp
	this.Size++
}

func (this *MyLinkedList) AddAtTail(val int) {
	cur := this.DummyTail
	temp := &Node{val, nil, nil}
	temp.Next = cur
	temp.Pre = cur.Pre
	cur.Pre.Next = temp
	cur.Pre = temp
	this.Size++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.Size {
		return
	}
	temp := &Node{Val: val}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	temp.Next = cur.Next
	temp.Pre = cur
	cur.Next.Pre = temp.Next
	cur.Next = temp
	this.Size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.Size {
		return
	}
	cur := this.DummyHead
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next.Next.Pre = cur
	cur.Next = cur.Next.Next
	this.Size--
}

func main() {
	list := Constructor() // 初始化链表
	list.AddAtHead(5)     // 在头部添加元素
	list.AddAtIndex(1, 2) // 在指定位置添加元素
	list.Get(1)
	list.AddAtHead(6) // 在头部添加元素
	list.AddAtTail(2) // 在尾部添加元素
	list.Get(3)
	list.AddAtHead(1) // 在头部添加元素

	//list.DeleteAtIndex(1)
	//list.Get(0)

}
