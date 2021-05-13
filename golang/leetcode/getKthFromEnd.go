package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func GetKthFromEnd(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	// 链表长度
	l := CountFromEnd(head)
	if l < k {
		return nil
	}
	for i := 0; i < l-k; i++ {
		head = head.Next
	}

	return head
}

func CountFromEnd(head *ListNode) int {
	if head == nil {
		return 0
	}
	return CountFromEnd(head.Next) + 1
}
