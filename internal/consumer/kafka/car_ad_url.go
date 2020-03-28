package kafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
)

func GetAdURL() car_ad.UrlChan {
	uc := make(car_ad.UrlChan)

	go func() {
		defer close(uc)
		c, err := kafka.NewConsumer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
			"group.id":          "ad_url_scraper_group",
			"auto.offset.reset": "earliest",
		})

		if err != nil {
			panic(err)
		}

		err = c.SubscribeTopics([]string{"autoreflex-ads"}, nil)
		if err != nil {
			panic(err)
		}

		for {
			msg, err := c.ReadMessage(-1)
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))

			if err != nil {
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, msg)
			}

			uc <- string(msg.Value)
		}

		c.Close()
	}()

	return uc
}
