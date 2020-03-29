package car_ad

import "context"

type Repository interface {
	Save(ctx context.Context, ads Chan)
}
