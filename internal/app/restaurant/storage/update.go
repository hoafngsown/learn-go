package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
)

func (s *sqlStore) UpdateRestaurant(ctx context.Context, id uint, data *restaurantmodel.RestaurantUpdate) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", id).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
