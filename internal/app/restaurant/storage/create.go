package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
)

func (s *sqlStore) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) *common.AppError {
	if err := s.db.Create(&restaurant).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
