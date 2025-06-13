package restaurantbiz

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) error {
	if err := biz.store.CreateRestaurant(ctx, restaurant); err != nil {
		return err
	}

	return nil
}
