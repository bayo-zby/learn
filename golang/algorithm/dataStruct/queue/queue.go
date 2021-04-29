package queue

// 队列，FIFO,单链
type Element struct {
	prev, next *Element    // 下一节
	value      interface{} // 空接口
}

type Queue struct {
	head, tail *Element //元素
	len        int      // 队列长度
}

/*入队
Desc:向队伍头部插入元素
In:
	@object q *quque,操作队列
	@params v interface{},插入值,可接受任意值
Out:
	无返回，nil队列使用Push会报错
*/
func (q *Queue) Push(v interface{}) {

	// 头插
	q.head = &Element{value: v, next: q.head}

	// 初始化插入时，不需要做双链对接
	if q.head.next != nil {
		// 第二节点的前一节指针链接到第一节
		q.head.next.prev = q.head
	}

	// 添加第一个元素时，即是头，也是尾
	if q.tail == nil {
		q.tail = q.head
	}

	// 添加长度
	q.len++

}

/*出队
Desc:从队伍尾部弹出元素
In:
	@object q *quque,操作对象
Out:
	无返回，空队列使用Pop会返回nil
*/
func (q *Queue) Pop() (res interface{}) {
	if q.tail == nil {
		return
	}
	res = q.tail.value
	q.tail = q.tail.prev // 尾指针移位
	if q.tail != nil {
		q.tail.next = nil // 避免内存泄漏
	}
	q.len--
	return
}

/*返回队列长度
Desc:返回队列长度
In:
	@object q *quque,操作对象
Out:
	int,返回队伍长度
*/
func (q *Queue) Len() int {
	return q.len
}
