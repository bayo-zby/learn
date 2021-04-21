package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.junmeitu.com"),
		colly.AllowURLRevisit(),
		colly.Async(true),
		colly.UserAgent(`Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4464.0 Safari/537.36 Edg/91.0.852.0`),
	)

	// 随机代理
	if p, err := proxy.RoundRobinProxySwitcher(
		"http://127.0.0.1:8080",
	); err == nil {
		c.SetProxyFunc(p)
	}

	c.OnHTML("ul.clearfix", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			link := h.ChildAttr("a", "href")
			title := h.ChildText("p")
			fmt.Println(link, title)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.junmeitu.com/beauty/")
	c.Wait()

}
