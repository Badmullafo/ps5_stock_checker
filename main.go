package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnHTML("span.a-size-medium.a-color-price", func(e *colly.HTMLElement) {
		// Goquery selection of the HTMLElement is in e.DOM
		//goquerySelection := e.DOM

		// Example Goquery usage
		//fmt.Println(e)
		ch := e.DOM.Text()
		fmt.Printf("test:%s", ch)
	})

	c.Visit("https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_3?dchild=1&keywords=ps5&qid=1630621564&sr=8-3")

	//c.Visit("https://www.amazon.co.uk/Xbox-RRT-00007-Series-X/dp/B08H93GKNJ")

}
