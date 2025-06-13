package restaurantmodel

import "learn-go/v2/internal/common"

const EntityName = "Restaurant"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Address         string `json:"address" gorm:"column:address"`
}

func (Restaurant) TableName() string { return "restaurants" }

func (r *Restaurant) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DBTypeRestaurant)
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Address         string `json:"address" gorm:"column:address"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantCreate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DBTypeRestaurant)
}

type RestaurantUpdate struct {
	common.SQLModel `json:",inline"`
	Name            string `json:"name" gorm:"column:name"`
	Address         string `json:"address" gorm:"column:address"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantUpdate) Mask(isAdminOrOwner bool) {
	r.GenUID(common.DBTypeRestaurant)
}
