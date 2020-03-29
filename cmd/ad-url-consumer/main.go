package main

import (
	"context"

	"github.com/rmasclef/autoreflex_scraper/internal/consumer/kafka"
	"github.com/rmasclef/autoreflex_scraper/internal/store/mongodb"
	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
)

// this program will get all car ad URLs from kafka
// then scrap the car ad page
// and save the car ad information into MongoDB
func main() {
	r := getAdRepository(context.TODO())

	urls := kafka.GetAdURLs()
	ads := car_ad.CollectAds(urls)

	r.Save(context.TODO(), ads)
}

func getAdRepository(ctx context.Context) car_ad.Repository {
	c := mongodb.NewClient(ctx)
	db := c.Database("autoreflex")

	return mongodb.NewCarAdRepository(db, 500)
}
