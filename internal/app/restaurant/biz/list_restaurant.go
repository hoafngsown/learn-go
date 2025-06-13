package restaurantbiz

import (
	"context"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/constants"
)

type ListRestaurantStore interface {
	ListWithCondition(ctx context.Context, paging *common.Paging, filter *restaurantmodel.Filter) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context, paging *common.Paging, filter *restaurantmodel.Filter) ([]restaurantmodel.Restaurant, error) {
	filter.Status = []constants.RestaurantStatus{constants.RestaurantStatusActive}

	result, err := biz.store.ListWithCondition(ctx, paging, filter)

	if err != nil {
		return nil, common.ErrorCannotListEntity(restaurantmodel.EntityName, err)
	}

	return result, nil
}
