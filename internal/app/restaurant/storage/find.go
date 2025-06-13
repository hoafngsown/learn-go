package restaurantstorage

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"

	"gorm.io/gorm"
)

func (s *sqlStore) FindWithCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName())

	var result restaurantmodel.Restaurant

	if err := db.Where(condition).First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrorDB(err)
	}

	return &result, nil
}
