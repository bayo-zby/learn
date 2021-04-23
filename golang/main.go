package main

import (
	"fmt"
	"learn/golang/algorithm/search"
)

func main() {
	res := search.BinarySearchRecursion([]int{1, 2, 3, 4, 5}, 0, 4, 1)
	fmt.Println(res)
}
