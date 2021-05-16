package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	c, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false), //debug使用
		)...,
	)

	ctx, cancel := chromedp.NewContext(c)
	defer cancel()

	// 参数编码
	// queryUrl := url.Values{}
	// queryUrl.Add("q", `拖鞋`)
	// queryUrl.Add("s", `1`)
	// searchUrl := fmt.Sprintf(`https://s.taobao.com/search?%v`, queryUrl.Encode())
	searchUrl := "http://www.customs.gov.cn/customs/302249/302266/302267/index.html"

	if err := chromedp.Run(ctx,
		chromedp.Navigate(searchUrl),
		// 等待页面
		chromedp.WaitVisible(`#easysite-page-wrap`),
	); err != nil {
		log.Fatal(err)
	}

}

func visitWeb(url string) chromedp.Tasks {
	return chromedp.Tasks{}
}
