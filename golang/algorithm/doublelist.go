package algorithm

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
	dln.Head = &Node{Key: k, Value: v, Next: dln.Head}
	dln.Size++
}

// 尾部添加节点
func (dln *DoubleListNode) PushTail(k string, v int) {
	dln.Tail = &Node{Key: k, Value: v, Prev: dln.Tail}
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
// 用索引位置来定位
func (dln *DoubleListNode) PushIndex(i int) {

}

// 弹出头部节点

// 弹出尾部节点
