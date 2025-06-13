package restaurantmodel

type Restaurant struct {
	ID      uint   `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
	Status  string `json:"status" gorm:"column:status"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	Id      uint   `json:"id" gorm:"column:id"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
