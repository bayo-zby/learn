package JD

import (
	"context"
	"fmt"

	"github.com/chromedp/chromedp"
)

func search(query string) string {
	c, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	if err := c.Run(tasks); err != nil {
		fmt.Println(err)
	}
}

func tasks() chromedp.tasks {

}
