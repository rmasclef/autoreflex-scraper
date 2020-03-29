package main

import (
	"github.com/rmasclef/autoreflex_scraper/internal/producer/kafka"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_brand"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_list"
)

// this program will scrap all list pages for each available car brand on http://www.autoreflex.com website
// it will then send all the car ad urls into a kafka topic
func main()  {
	// get brands
	carBrands := car_brand.Collect()
	// get pages URLs (pagination) for each brand
	pageUrls := car_list.Collect(carBrands)
	// get all ad "standalone page" URLs
	adUrls := car_ad.CollectURLs(pageUrls)

	kafka.SendURL(adUrls)
}
