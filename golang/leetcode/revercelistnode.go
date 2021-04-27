package leetcode

import "learn/golang/algorithm/dataStruct"

/*
 * @brief: 迭代法反转链表
 * @param: l 原链表
 * @return: 反转链表
 */
func ReverseListNode(l *dataStruct.ListNode) *dataStruct.ListNode {
	var prev, next *dataStruct.ListNode
	for l.Next != nil {
		//
		prev,next = l,l.Next
		l.Next = nil
		l = next
		l.Next =
	}

	return ReverseListNode(l)

}
