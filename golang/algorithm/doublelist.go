package algorithm

import "fmt"

type Node struct {
	Key   string // 键
	Value int    // 值
	Prev  *Node  // 前一节点
	Next  *Node  //
}

type DoubleListNode struct {
	Head     *Node // 头节点
	Tail     *Node // 尾节点
	Capacity int   // 最大容量
	Size     int   // 链表大小
}

// 新建双向链表
func NewDouleListNode() *DoubleListNode {
	// 默认容量为0xffff
	return &DoubleListNode{Capacity: 0xffff}
}

// 头部添加节点
func (dln *DoubleListNode) PushHead(k string, v int) {
	head := &Node{Key: k, Value: v, Next: dln.Head}
	if dln.Head != nil {
		dln.Head.Prev = head
	}
	dln.Head = head

	if dln.Tail == nil {
		dln.Tail = dln.Head
	}

	dln.Size++
}

// 尾部添加节点
func (dln *DoubleListNode) PushTail(k string, v int) {
	tail := &Node{Key: k, Value: v, Prev: dln.Tail}
	if dln.Tail != nil {
		dln.Tail.Next = tail
	}
	dln.Tail = tail

	if dln.Head == nil {
		dln.Head = dln.Tail
	}
	dln.Size++
}

// 删除任意节点
func (dln *DoubleListNode) DeleteKey(k string) {
	next := dln.Head
	for next != nil {
		if next.Key == k {
			next.Prev.Next, next.Next.Prev = next.Next, next.Prev
			// 清空指针，由GC处理释放内存
			next.Prev = nil
			next.Next = nil
			dln.Size--
			break
		}
	}
}

// 增加任意位置的节点
// func (dln *DoubleListNode) PushIndex(i int)

// 弹出头部节点
func (dln *DoubleListNode) PopHead() *Node {
	if dln.Head == nil {
		return nil
	}
	head := dln.Head
	dln.Head = dln.Head.Next
	dln.Head.Prev = nil
	dln.Size--
	return head
}

// 弹出尾部节点
func (dln *DoubleListNode) PopTail() *Node {
	if dln.Tail == nil {
		return nil
	}
	tail := dln.Tail
	dln.Tail = dln.Tail.Prev
	// 清理指针
	dln.Tail.Next = nil
	dln.Size--
	return tail
}

// 展示链表内容
func Print(dln *DoubleListNode) {
	p := dln.Head
	for p != nil {
		fmt.Printf("%p:%+v\n", p, p)
		p = p.Next
	}
}
