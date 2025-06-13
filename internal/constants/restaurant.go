package constants

type RestaurantStatus int

type RestaurantStatusStr string

const (
	RestaurantStatusActive   RestaurantStatus = 1
	RestaurantStatusInactive RestaurantStatus = 0
)

const (
	RestaurantStatusActiveEnum   RestaurantStatusStr = "active"
	RestaurantStatusInactiveEnum RestaurantStatusStr = "inactive"
)

func (rs RestaurantStatus) ToString() RestaurantStatusStr {
	switch rs {
	case RestaurantStatusActive:
		return RestaurantStatusActiveEnum
	case RestaurantStatusInactive:
		return RestaurantStatusInactiveEnum
	}

	return RestaurantStatusActiveEnum
}
