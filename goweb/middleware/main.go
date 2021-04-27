package main

import (
	"fmt"
	"learn/goweb/middleware/simhash"
)

func main() {
	words := simhash.WordWeight("我感觉知乎有个非常坑的地方，就是每个人都说Arch如何如何有门槛..")
	fmt.Println(words)
}
