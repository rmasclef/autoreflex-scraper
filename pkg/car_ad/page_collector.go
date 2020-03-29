package car_ad

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

const (
	brand int = iota
	model
)

const (
	kilometer int = iota
	modelYear
	location
	energy
	transmission
	power
)

func CollectAds(urls URLChan) Chan {
	var ac = make(Chan, 20000)

	go func() {
		// close the chan when all the car ads have been treated
		defer close(ac)

		for url := range urls {

			// create new car ad
			ad := &Ad{}
			c := getCollector(ad)

			err := c.Visit("http://www.autoreflex.com/" + url)
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

	// Limit the maximum parallelism to 10
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
	// get Brand, Model
	c.OnHTML("div.header h1", func(elt *colly.HTMLElement) {
		elt.ForEach("a", func(i int, liElt *colly.HTMLElement) {
			switch i {
			case brand:
				ad.Brand = liElt.Text
				break
			case model:
				ad.Model = liElt.Text
				break
			default:
				log.Printf("unrecognised header (brand, model) %s", elt.Text)
			}
		})
		// @TODO find a way to get vehicle information (i.e information that is after the brand and model)
		// fmt.Println("-------------------------------------------------------> spec: ", elt.ChildText("*:not(a)"))
	})
	// get Price
	c.OnHTML("div.prix", func(elt *colly.HTMLElement) {
		// @TODO extract currency from value
		ad.Price = elt.Text
	})
	// get specifications
	c.OnHTML("div.specs ul", func(elt *colly.HTMLElement) {
		elt.ForEach("li", func(i int, liElt *colly.HTMLElement) {
			switch i {
			case kilometer:
				// @TODO extract value from "km"
				ad.Kilometers = liElt.Text
				break
			case modelYear:
				ad.ModelYear = liElt.Text
				break
			case location:
				ad.Location = liElt.Text
				break
			case energy:
				ad.Energy = liElt.Text
				break
			case transmission:
				ad.Transmission = liElt.Text
				break
			case power:
				// @TODO extract HP from fiscal power
				ad.Power = liElt.Text
				break
			default:
				log.Printf("unrecognised spec %s", liElt.Text)
			}
		})
	})

	return c
}
