package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/constants"
)

type DeleteRestaurantStore interface {
	FindWithCondition(ctx context.Context, condition map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error)
	DeleteRestaurant(ctx context.Context, id uint) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(ctx context.Context, id uint) error {
	existRestaurant, err := biz.store.FindWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrorCannotGetEntity(restaurantmodel.EntityName, err)
	}

	if existRestaurant.Status == string(constants.RestaurantStatusInactiveEnum) {
		return common.ErrorCannotDeleteEntity(restaurantmodel.EntityName, ErrRestaurantDeleted)
	}

	err = biz.store.DeleteRestaurant(ctx, id)

	if err != nil {
		return common.ErrorCannotDeleteEntity(restaurantmodel.EntityName, err)
	}

	return nil
}

var (
	ErrRestaurantDeleted = errors.New("cannot delete restaurant")
)
