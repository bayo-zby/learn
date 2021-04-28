package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintLn(l *ListNode) {
	for l != nil {
		fmt.Printf("%+v\n", l)
		l = l.Next
	}
}

func NewListNode(sInt []int) *ListNode {
	// 哨兵
	head := &ListNode{}
	curr := head
	for _, v := range sInt {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return head.Next
}
