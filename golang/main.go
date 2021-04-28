package main

import (
	"fmt"
	"learn/golang/leetcode"
	"unsafe"
)

func main() {
	res := leetcode.CountPrimeEratothenes(5)
	fmt.Println(res)
	fmt.Println(unsafe.Sizeof(1))
}
