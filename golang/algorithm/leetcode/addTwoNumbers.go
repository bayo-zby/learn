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
	head := &ListNode{}
	recursionAdd(l1, l2, head)
	<-ch
	return head
}

func recursionAdd(l1 *ListNode, l2 *ListNode, tail *ListNode) {
	// 考虑进位导致链表扩容的情况
	if l1 == nil && l2 == nil {
		// 清掉尾环
		tail = nil
		ch <- struct{}{}
		return
	}

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
	tail.Next = &ListNode{Val: sum}

	//递归传递
	go recursionAdd(l1, l2, tail.Next)
}
