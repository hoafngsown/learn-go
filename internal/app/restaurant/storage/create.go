package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) error {
	return s.db.Create(restaurant).Error
}
