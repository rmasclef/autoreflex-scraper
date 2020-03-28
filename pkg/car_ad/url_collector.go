package car_ad

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func ExtractPageUrls(puc PaginationURLChan) UrlChan {
	var err error

	urlChan := make(UrlChan, 10000000)

	go func() {
		defer close(urlChan)

		c := getPageUrlsCollector()

		// get all available car ads on the current page
		c.OnHTML("tr[star-id]>td>h2>a", func(elt *colly.HTMLElement) {
			adUrl := elt.Attr("href")
			fmt.Printf("ad url found : %s\n", adUrl)
			urlChan <- adUrl
		})

		for pageUrl := range puc {
			// scrape the current brandID/pageNumber page (containing a list of ad URLs
			err = c.Visit("http://www.autoreflex.com"+string(pageUrl))
		}
		c.Wait()
	}()

	return urlChan
}

func getPageUrlsCollector() *colly.Collector {
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