package main

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

func main() {

	doc, err := htmlquery.LoadDoc(`C:\Users\viruser.v-desktop\Desktop\a.html`)
	if err != nil {
		panic(`valid document`)
	}

	list := htmlquery.Find(doc, `//div[@id='etfs']//table/tbody/tr/td[1]`)
	if err != nil {
		panic(`not a valid XPATH expression`)
	}

	for _, n := range list {
		// 展示字符串
		fmt.Println(htmlquery.InnerText(n)) // output @href value
		childList := htmlquery.FindOne(n, `a`)
		fmt.Println(htmlquery.SelectAttr(childList, `href`))
	}
}
