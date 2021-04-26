package main

import "fmt"

func main() {
	a := 0b1101
	fmt.Printf("%08b\n", a)
	fmt.Printf("%08b\n", a^a)
}
