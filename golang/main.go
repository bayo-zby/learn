package main

import (
	"fmt"
	"gitee/learn/golang/algorithm"
)

func main() {
	dln := algorithm.NewDouleListNode()
	for i := 0; i < 10; i++ {
		dln.PushHead(fmt.Sprintf("%d", i), i)
	}
	algorithm.Print(dln)
}
