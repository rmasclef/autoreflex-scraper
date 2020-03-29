package car_ad

type URLChan chan string

type Chan chan Ad

type Ad struct {
	Location     string `bson:"location,omitempty"`
	Brand        string `bson:"brand,omitempty"`
	Model        string `bson:"model,omitempty"`
	ModelYear    string `bson:"model_year,omitempty"`
	Kilometers   string `bson:"kilometers,omitempty"`
	Energy       string `bson:"energy,omitempty"`
	Transmission string `bson:"transmission,omitempty"`
	// @TODO extract HP from fiscal power
	Power string `bson:"power,omitempty"`
	// @TODO extract currency from value
	Price  string  `bson:"price,omitempty"`
	Images []Image `bson:"images,omitempty"`
}

type Image struct {
	URL    string `bson:"url,omitempty"`
	IsMain bool   `bson:"is_main,omitempty"`
}

// type Price struct {
// 	Value float32 `bson:"value"`
// 	Currency string `bson:"currency"`
// }
