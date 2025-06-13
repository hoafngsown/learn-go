package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/constants"
)

func (s *sqlStore) DeleteRestaurant(ctx context.Context, id uint) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", id).Updates(restaurantmodel.Restaurant{
		Status: string(constants.RestaurantStatusInactiveEnum),
	}).Error; err != nil {
		return err
	}

	return nil
}
