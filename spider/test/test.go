package main

import (
	"fmt"
	"net/url"
)

func main() {
	queryUrl := url.Values{}
	queryUrl.Add("q", `拖鞋`)
	fmt.Println(queryUrl.Encode())
}
