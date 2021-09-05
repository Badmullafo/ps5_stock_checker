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

	c.OnHTML("div[id=price]", func(e *colly.HTMLElement) {

		price := e.ChildText(".priceBlockBuyingPriceString")

		// You don't need the whole path

		if price != "" {
			fmt.Printf("price:%s\n", price)
		} else {
			fmt.Println("Could not find price - out of stock")
		}

		//fmt.Printf("test:%s\n", e.Text)

		//fmt.Printf("test:%s\n", e.ChildText("a-size-medium.a-color.price"))
		//e.ForEach("a-size-medium.a-color.price", func(_ int, el *colly.HTMLElement) {
		//	fmt.Println(el.ChildText())
		//})

	})

	//c.Visit("https://www.amazon.co.uk/PlayStation-9395003-5-Console/dp/B08H95Y452/ref=sr_1_3?dchild=1&keywords=ps5&qid=1630621564&sr=8-3")

	c.Visit("https://www.amazon.co.uk/Xbox-RRT-00007-Series-X/dp/B08H93GKNJ")
	c.Visit("https://www.amazon.co.uk/Konsole-Sony-Playstation-D-Chasis-Controller/dp/B01LQF9RDI/ref=sr_1_1?dchild=1&keywords=ps4&qid=1630872165&sr=8-1")

}
