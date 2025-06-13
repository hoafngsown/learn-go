package restaurantbiz

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
)

type CreateRestaurantStore interface {
	CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) *common.AppError
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, restaurant *restaurantmodel.RestaurantCreate) *common.AppError {
	if err := biz.store.CreateRestaurant(ctx, restaurant); err != nil {
		return common.ErrorCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
