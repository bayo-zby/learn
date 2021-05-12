package main

import (
	"context"
	"fmt"
	"log"

	"github.com/chromedp/chromedp"
)

const HOST = "http://www.fullgoal.com.cn/PCF/html/PCF21/index.html?fundCode=517101&fundcode=517100"

func main() {
	// 创建模拟浏览器对象
	ctx, canel := chromedp.NewContext(
		context.Background(),
		// 模拟浏览器覆盖日志
		chromedp.WithLogf(log.Printf),
	)
	defer canel()

	var pageContent string
	//
	url := HOST + ``

	// 执行内容
	if err := chromedp.Run(ctx, *tasks(url, &pageContent)...); err != nil {
		log.Fatal(err)
	}

	fmt.Println(pageContent)
}

/**
 *	@brief	执行具体执行操作
 *	@params url,起始URL链接
 *	@params contentReader,正文指针
 *	@return 操作集合
 */
func tasks(url string, contentReader *string) *chromedp.Tasks {
	return &chromedp.Tasks{
		chromedp.Navigate(url),
		// 等待到页面元素出现
		chromedp.WaitVisible(``, chromedp.ByID),
		// 输出页面内容
		chromedp.OuterHTML(`.//body`, contentReader),
		// chromedp.Nodes(`.//a[@class="titlelnk"]`, &nodes),
	}
}
