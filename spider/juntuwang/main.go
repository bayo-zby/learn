package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c1 := colly.NewCollector(
		colly.AllowedDomains("www.junmeitu.com"),
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.UserAgent(`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4464.0 Safari/537.36 Edg/91.0.852.0`),
	)

	// 随机代理
	// if p, err := proxy.RoundRobinProxySwitcher(
	// 	"http://127.0.0.1:8080",
	// ); err == nil {
	// 	c1.SetProxyFunc(p)
	// }

	// 克隆一个采集器
	c2 := c1.Clone()

	// 列表内容采集
	c1.OnHTML(`div.pic-list > ul.clearfix`, func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(item.ChildAttr("a", "href"))
			title := item.ChildText("a > p")
			fmt.Println(link, title)
			ctx := colly.NewContext()
			ctx.Put("link", link)
			ctx.Put("title", title)
			// 传递数据给c2
			c2.Request("GET", link, nil, ctx, nil)
		})
		// 翻页
	})

	// 文章内容采集
	c2.OnHTML(`div.content`, func(e *colly.HTMLElement) {
		imgUrl := e.ChildAttr(`div.pictures img`, "href")
		fmt.Println(imgUrl)
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c1.Visit("https://www.junmeitu.com/model/xunuo.html"); err != nil {
		fmt.Println(err.Error())
	}
	c1.Wait()
}
