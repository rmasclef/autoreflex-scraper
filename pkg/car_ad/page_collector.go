package car_ad

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gocolly/colly"

	"github.com/rmasclef/autoreflex_scraper/pkg/car_brand"
)

func ExtractPages(brandChan car_brand.Chan) PageURLChan {
	pg := make(PageURLChan, 1000)

	go func() {
		defer close(pg)

		// for all brands
		for brand := range brandChan {
			c := getNbPagesCollector()

			// get number of pages
			c.OnHTML("ul.pagination li:nth-last-child(2)", func(elt *colly.HTMLElement) {
				nbPages, _ := strconv.Atoi(elt.Text)
				for pageNumber := 1; pageNumber <= nbPages; pageNumber++ {
					// send page list url to be scraped
					pg <- PageURL(fmt.Sprintf(url, brand.ID, pageNumber))
				}
			})

			// we scrap the first brand ad page in order to get the number of available pages
			err := c.Visit(fmt.Sprintf("http://www.autoreflex.com"+url, brand.ID, 1))
			if err != nil {
				panic(err)
			}

			c.Wait()
		}
	}()

	return pg
}

func getNbPagesCollector() *colly.Collector {
	c := colly.NewCollector(
		colly.AllowedDomains("www.autoreflex.com"),
		colly.Async(true),
	)

	c.AllowURLRevisit = false

	// Limit the maximum parallelism to 2
	// This is necessary if the goroutines are dynamically
	// created to control the limit of simultaneous requests.
	//
	// Parallelism can be controlled also by spawning fixed
	// number of go routines.
	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	if err != nil {
		panic(err)
	}

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
