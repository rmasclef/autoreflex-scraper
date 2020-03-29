package car_ad

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

func CollectAds(urls URLChan) Chan {
	var ac = make(Chan, 20000)

	go func() {
		// close the chan when all the car ads have been treated
		defer close(ac)

		for url := range urls{

			// create new car ad
			ad := &Ad{}
			c := getCollector(ad)

			err := c.Visit("http://www.autoreflex.com/"+url)
			if err != nil && err != colly.ErrAlreadyVisited {
				panic(err)
			}
			// the collector will fill the ad object
			c.Wait()

			// once the collect is finished, we send the ad into ad chan
			ac <- *ad
		}
	}()

	return ac
}

func getCollector(ad *Ad) *colly.Collector {
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
	err := c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 10})
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

	// get Price
	c.OnHTML("div.prix", func(elt *colly.HTMLElement) {
		log.Println("Price found", elt.Text)
		// @TODO extract currecy from value
		ad.Price = elt.Text
	})
	// // get Modele
	// c.OnHTML("", func(elt *colly.HTMLElement) {
	//
	// })
	// // get description
	// c.OnHTML("", func(elt *colly.HTMLElement) {
	//
	// })

	return c
}
