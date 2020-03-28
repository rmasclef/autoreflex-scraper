package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/gocolly/colly"

	"github.com/rmasclef/autoreflex_scraper/internal/consumer/kafka"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
)

// this program will get all car ad URLs from kafka
// then scrap the car ad page
// and save the car ad information into MongoDB
func main() {
	urls := kafka.GetAdURL()

	var wg sync.WaitGroup
	var ac = make(chan car_ad.Ad, 2000)

	go func() {
		for url := range urls{
			wg.Add(1)
			go func() {
				defer wg.Done()

				// create new car ad
				ad := &car_ad.Ad{}
				c := getCollector(ad)

				err := c.Visit("http://www.autoreflex.com/"+url)
				if err != nil {
					panic(err)
				}
				c.Wait()
				fmt.Println("send Ad into chan")
				ac <- *ad
			}()
		}
	}()

	for ad := range ac {
		fmt.Println("SAVE ADD TO MONGO ------------>>>>>>>>>>>>>>"+ ad.Price)
	}

	wg.Wait()
}

func getCollector(ad *car_ad.Ad) *colly.Collector {
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
