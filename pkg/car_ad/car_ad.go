package car_ad

const url = "/%s.0.-1.-1.-1.0.999999.1900.999999.-1.99.0.%d"

type PageURLChan chan PageURL
type PageURL string

type UrlChan chan string

type Ad struct {
	Brand string `bson:"brand"`
	Model string `bson:"brand"`
	// @TODO extract currecy from value
	Price string `bson:"price"`
	Images []Image `bson:"images"`

	// @TODO add other information
}

type Image struct {
	URL string `bson:"url"`
	IsMain bool `bson:"is_main"`
}

// type Price struct {
// 	Value float32 `bson:"value"`
// 	Currency string `bson:"currency"`
// }
