package leetcode

import (
	"fmt"
)

/*
	该文件用于存储leetcode的所有数据结构，以及数据结构可能用到的方法
*/

// 链表
/*
	structName:ListNode
	description:头插单链表，无哨兵
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*	打印展示链表所有节点
	funcName :PrintListNode
*/
func PrintListNode(ln *ListNode) {
	for ln != nil {
		fmt.Println(ln.Val)
		ln = ln.Next
	}
}

func NewListNode(sInt []int) *ListNode {
	head := &ListNode{}
	ln := head
	for _, v := range sInt {
		ln.Val = v
		ln.Next = &ListNode{}
		ln = ln.Next
	}
	ln.Next = nil
	return head
}
