package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
)

func main() {
	c := colly.NewCollector(colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36"),
		colly.MaxDepth(1),
		colly.Debugger(&debug.LogDebugger{}))

	c.OnHTML("ul[class='note-list']", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, item *colly.HTMLElement) {
			href := item.ChildAttr("div[class='content'] > a[class='title']", "href")
			title := item.ChildText("div[class='content'] > a[class='title']")
			summary := item.ChildText("div[class='content'] > p[class='abstract']")
			fmt.Println(title, href)
			fmt.Println(summary)
			fmt.Println()
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	err := c.Visit("https://www.jianshu.com")
	if err != nil {
		fmt.Println(err.Error())
	}
}
