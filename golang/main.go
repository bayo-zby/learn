package main

import (
	"fmt"
	"learn/goweb/middleware/simhash"
)

func main() {
	str := simhash.WordWeight("我来到北京清华大学")
	fmt.Println(str)
}
