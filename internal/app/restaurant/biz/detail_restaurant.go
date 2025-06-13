package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/constants"
)

type DetailRestaurantStore interface {
	FindWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
}

type detailRestaurantBiz struct {
	store DetailRestaurantStore
}

func NewDetailRestaurantBiz(store DetailRestaurantStore) *detailRestaurantBiz {
	return &detailRestaurantBiz{store: store}
}

func (biz *detailRestaurantBiz) DetailRestaurant(ctx context.Context, id uint) (*restaurantmodel.Restaurant, error) {
	status := string(constants.RestaurantStatusActiveEnum)
	restaurant, err := biz.store.FindWithCondition(
		ctx,
		map[string]interface{}{
			"id":     id,
			"status": status,
		},
	)

	if err != nil {
		return nil, common.ErrorCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if restaurant == nil {
		return nil, common.ErrorNotFound(ErrRestaurantNotFound)
	}

	return restaurant, nil
}

var (
	ErrRestaurantNotFound = errors.New("restaurant not found")
)
