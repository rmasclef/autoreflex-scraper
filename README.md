# autoreflex-scraper

A scraper that will get all car ads from autoreflex.com

## Quick tour

This system is composed of two binaries:
- ad-url-producer
- ad-url-consumer

### ad-url-producer

this binary will perform the following steps on autoreflex.com :

1. get all available brands
2. for each brands -> get the pagination URLs 
3. for each pagination urls -> get the Ad URL
4. for each Ad URL -> send them into a kafka topic

### ad-url-consumer

this binary will perform the following steps :

1. consume the messages produced by `ad-url-producer` from the kafka topic
2. for each messages (i.e Ad URL) -> scrap the Ad information (still work to do here) -> transform them into an Ad model
3. for each generated Ad -> save it into a Mongo databse (not present yet)

## Make it work

### DEV

#### start a Kafka server
`docker-compose up -d`

#### start a consumer
`go run ./cmd/ad-url-consumer`

#### start a producer
`go run ./cmd/ad-url-producer`

> NOTE Docker version is not ready yet for the consumer & producer ...
