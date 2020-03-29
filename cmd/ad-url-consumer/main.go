package main

import (
	"fmt"

	"github.com/rmasclef/autoreflex_scraper/internal/consumer/kafka"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
)

// this program will get all car ad URLs from kafka
// then scrap the car ad page
// and save the car ad information into MongoDB
func main() {
	urls := kafka.GetAdURLs()
	ads := car_ad.CollectAds(urls)

	for ad := range ads {
		fmt.Println("SAVE TO MONGO ------------>>>>>>>>>>>>>>"+ ad.Price)
	}
}
