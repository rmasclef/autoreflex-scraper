package car_brand

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func Collect() Chan {
	bc := make(Chan, 30)

	go func() {
		defer close(bc)

		c := getCollector()

		// get car brand list
		c.OnHTML(`select[id=marque_home]`, func(e *colly.HTMLElement) {
			e.ForEach("option", func(_ int, el *colly.HTMLElement) {
				if el.Attr("value") == "0" {
					// we don't want to get select options that do not match any car brand
					return
				}

				// send the car brand into our carBrand chan if it does not exists yet
				bc <- Brand{
					ID:   el.Attr("value"),
					Name: el.Text,
				}
			})
		})

		err := c.Visit("http://www.autoreflex.com")
		if err != nil && err != colly.ErrAlreadyVisited {
			panic(err)
		}
	}()

	return bc
}

// @TODO this collector is the same as car_ad ones -> make a factory or something like this
func getCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.autoreflex.com"),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Visited", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})

	return c
}
