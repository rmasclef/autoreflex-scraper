package car_ad

type URLChan chan string

type Chan chan Ad

type Ad struct {
	Brand string `bson:"brand,omitempty"`
	Model string `bson:"model,omitempty"`
	// @TODO extract currency from value
	Price string `bson:"price,omitempty"`
	Images []Image `bson:"images,omitempty"`

	// @TODO add other information
}

type Image struct {
	URL string `bson:"url,omitempty"`
	IsMain bool `bson:"is_main,omitempty"`
}

// type Price struct {
// 	Value float32 `bson:"value"`
// 	Currency string `bson:"currency"`
// }
