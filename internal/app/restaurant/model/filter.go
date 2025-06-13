package restaurantmodel

import "learn-go/v2/internal/constants"

type Filter struct {
	OwnerId int                             `json:"owner_id,omitempty" form:"owner_id"`
	Status  []constants.RestaurantStatusStr `json:"status,omitempty" form:"status"`
}
