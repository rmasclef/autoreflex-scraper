package car_ad

type URLChan chan string

type Chan chan Ad

type Ad struct {
	Brand string `bson:"brand"`
	Model string `bson:"brand"`
	// @TODO extract currency from value
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
