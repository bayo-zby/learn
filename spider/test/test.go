package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const url = "http://www.i618.com.cn/gsyw/zcgl/qxcp/productSZ9122/zgcpjs.shtml"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}

	// req.Header.Add("User-Agent","curl/1.7")

	// t := &http.Transport{Proxy: ("http://hexin:hx300033@119.97.156.72:888")}

	client := &http.Client{
		// CheckRedirect: red,
		// Transport: t,
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
