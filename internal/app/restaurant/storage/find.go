package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
)

func (s *sqlStore) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	var result restaurantmodel.Restaurant

	if err := db.Where(condition).First(&result).Error; err != nil {
		return nil, common.ErrorDB(err)
	}

	return &result, nil
}
