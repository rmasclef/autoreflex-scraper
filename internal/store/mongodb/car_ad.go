package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rmasclef/autoreflex_scraper/pkg/car_ad"
)

const collectionName = "ad"

type CarAdRepository struct {
	col *mongo.Collection
	bulkSize int
}

func NewCarAdRepository(db *mongo.Database, bulkSize int) *CarAdRepository {
	return &CarAdRepository{
		col: db.Collection(collectionName),
		bulkSize: bulkSize,
	}
}

func (r *CarAdRepository) Save(ctx context.Context, ads car_ad.Chan) {
	var adBuffer []interface{}

	for ad := range ads {
		log.Println("adding new add into buffer")
		// add ad into buffer
		adBuffer = append(adBuffer, ad)
		log.Printf("buffer size %d\n bulk size %d\n", len(adBuffer), r.bulkSize)
		// bulk insert ads when the buffer is full
		if len(adBuffer) == r.bulkSize {
			log.Println("Bulk insert ads")
			r.insertMany(ctx, adBuffer)
			// flush buffer
			adBuffer = []interface{}{}
		}
	}
	// all ads have been treated
	// we make a last insert if there is still ads in the buffer
	if len(adBuffer) > 0 {
		log.Println("Last insert ads")
		r.insertMany(ctx, adBuffer)
		// flush buffer
		adBuffer = []interface{}{}
	}

	log.Println("All car ads have been saved into 'ad' collection")
}

func (r *CarAdRepository) insertMany(ctx context.Context, ads []interface{}) {
	_, err := r.col.InsertMany(ctx, ads)
	if err != nil {
		log.Fatal(err)
	}
}
