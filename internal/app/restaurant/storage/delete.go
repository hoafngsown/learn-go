package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/constants"
)

func (s *sqlStore) DeleteRestaurant(ctx context.Context, id uint) error {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	if err := db.Where("id = ?", id).Updates(map[string]interface{}{
		"status": string(constants.RestaurantStatusInactiveEnum),
	}).Error; err != nil {
		return common.ErrorDB(err)
	}

	return nil
}
