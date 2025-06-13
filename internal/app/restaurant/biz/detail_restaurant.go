package restaurantbiz

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
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
	restaurant, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, err
	}

	return restaurant, nil
}
