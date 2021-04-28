package leetcode

/*
 * @brief: 迭代法反转链表
 * @param: l 原链表
 * @return: 反转链表
 */
func ReverseListNodeIterate(l *ListNode) *ListNode {
	var head, next *ListNode
	for l != nil {
		// 后链转next
		next = l.Next
		// 后链转前链
		l.Next = head
		// 前链转当前链
		head = l
		// 转至下一个节点
		l = next
	}
	return head
}

/*
 * @brief: 递归法反转链表
 * @param: l,原链表
 * @ret: 反转链表
 */
func ReverseListNodeRecursion(l *ListNode) *ListNode {
	// 回溯条件
	// 考虑空链情况
	if l == nil || l.Next == nil {
		return l
	}
	// 下一个节点为空
	next := ReverseListNodeRecursion(l.Next)
	l.Next.Next = l
	// 该层断开的链条会下下一次回溯接上
	l.Next = nil
	return next
}
