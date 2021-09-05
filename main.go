package main

import (
	"fmt"

	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("div[id=availability]", func(e *colly.HTMLElement) {

		stock := strings.Split(e.ChildText("span.a-size-medium.a-color-price"), "\n")[0]
		fmt.Printf("test:%s\n", stock)

	})

	c.Visit("https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_3?dchild=1&keywords=ps5&qid=1630621564&sr=8-3")

	c.Visit("https://www.amazon.co.uk/Xbox-RRT-00007-Series-X/dp/B08H93GKNJ")
	c.Visit("https://www.amazon.co.uk/Konsole-Sony-Playstation-D-Chasis-Controller/dp/B01LQF9RDI/ref=sr_1_1?dchild=1&keywords=ps4&qid=1630872165&sr=8-1")

}
