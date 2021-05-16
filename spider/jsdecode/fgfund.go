package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)

	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	ctx, cancel = chromedp.NewContext(
		ctx,
		chromedp.WithLogf(log.Printf),
	)
	defer cancel()

	var url string = "http://www.fullgoal.com.cn/PCF/html/PCF21/index.html?fundCode=517101"

	// 页面内容
	// var htmlContent string
	if err := chromedp.Run(ctx, mytasks(url)); err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(htmlContent)
}

var htmlContent string

func mytasks(url string) chromedp.Tasks {
	return chromedp.Tasks{
		// 页面
		chromedp.Navigate(url),
		// 等待内容展示完成
		chromedp.WaitVisible(`#constituentStockInfoListInfo > table`),
		// 返回文本内容
		chromedp.OuterHTML(`//html`, &htmlContent),
	}
}
