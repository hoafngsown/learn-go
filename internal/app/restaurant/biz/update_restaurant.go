package restaurantbiz

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
)

type UpdateRestaurantStore interface {
	UpdateRestaurant(ctx context.Context, id uint, data *restaurantmodel.RestaurantUpdate) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantStore
}

func NewUpdateRestaurantBiz(store UpdateRestaurantStore) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id uint, data *restaurantmodel.RestaurantUpdate) error {
	if err := biz.store.UpdateRestaurant(ctx, id, data); err != nil {
		return err
	}

	return nil
}
