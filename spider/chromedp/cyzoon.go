package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const HOST = "https://data.cyzone.cn"

func main() {
	// 创建模拟浏览器对象
	ctx, canel := chromedp.NewContext(
		context.Background(),
		// 模拟浏览器覆盖日志
		chromedp.WithLogf(log.Printf),
	)
	defer canel()

	var nodes []*cdp.Node
	page := 1
	url := HOST + fmt.Sprintf(`/event/list-0-1-0-0-0-0-%d/0`, page)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		// 等待到页面元素出现
		chromedp.WaitVisible(``, chromedp.ByID),
		chromedp.Nodes(`.//a[@class="titlelnk"]`, &nodes),
	)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("get nodes", len(nodes))
	for _, node := range nodes {
		fmt.Println(node.Children[0].NodeValue, ":", node.AttributeValue("href"))
	}
}
