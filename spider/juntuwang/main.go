package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gocolly/colly"
)

type fileInf struct {
	title    string
	filename string
	imgUrl   string
}

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

			// 传递数据给c2
			c2.Visit(link)
		})
		// 翻页
	})

	// 文章内容采集
	c2.OnHTML(`.content`, func(e *colly.HTMLElement) {
		imgUrl := e.ChildAttr(`.pictures > .con_img`, "src")
		pageIndex := e.ChildText(`.pages > span`)
		filepath := e.ChildText(`.title`)
		saveimg(filepath, pageIndex, imgUrl)

		nextDom := e.DOM.Find(".pages > a").Last()
		if nextDom.Text() == "下一页" {
			if nextUrl := nextDom.AttrOr("href", ""); nextUrl != e.Request.URL.Path {
				c2.Visit(e.Request.AbsoluteURL(nextUrl))
			}

		}
	})

	c1.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c2.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	if err := c1.Visit("https://www.junmeitu.com/model/xunuo.html"); err != nil {
		fmt.Println(err.Error())
	}

	c1.Wait()
	c2.Wait()
}

// 图片保存
func saveimg(filepath, filename, imgUrl string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	filepath = "/run/media/bayo/DATA/DATA/PIC/" + filepath
	// 读取字节流
	req, err := http.NewRequest("GET", imgUrl, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36")
	req.Header.Set("referer", "https://www.junmeitu.com")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		panic(err)
	}

	// 创建文件夹
	_ = os.MkdirAll(filepath, os.ModePerm)
	lastIndex := strings.LastIndex(imgUrl, ".")
	out, err := os.Create(filepath + "/" + filename + imgUrl[lastIndex:len(imgUrl)-1])
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

}
