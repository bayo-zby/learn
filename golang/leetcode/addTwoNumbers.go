package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

var ch = make(chan struct{})

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 哨兵
	head := &ListNode{Val: 99}
	recursionAdd(l1, l2, head)
	<-ch
	return head.Next
}

func recursionAdd(l1 *ListNode, l2 *ListNode, tail *ListNode) {
	// 考虑进位导致链表扩容的情况
	if l1 == nil && l2 == nil {
		ch <- struct{}{}
		return
	}

	if tail.Next == nil {
		tail.Next = &ListNode{}
	}
	tail = tail.Next

	// 取值
	var n1, n2 int
	if l1 != nil {
		n1 = l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		n2 = l2.Val
		l2 = l2.Next
	}

	sum := n1 + n2 + tail.Val
	// 遗留数+当前循环值=当前累计值
	tail.Val, sum = sum%10, sum/10
	if sum != 0 {
		tail.Next = &ListNode{Val: sum}
	}

	//递归传递
	go recursionAdd(l1, l2, tail)
}
