package main

import (
	"gitee/learn/golang/algorithm/leetcode"
)

func main() {
	testList1 := []int{1, 3, 5, 7}
	testList2 := []int{1, 3, 5, 7}
	ln1 := leetcode.NewListNode(testList1)
	ln2 := leetcode.NewListNode(testList2)
	res := leetcode.AddTwoNumbers(ln1, ln2)
	leetcode.PrintListNode(res)
}
