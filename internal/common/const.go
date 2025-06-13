package common

import "learn-go/v2/internal/interfaces"

const DBTypeRestaurant = 1

func AppRecover(util *interfaces.Util) {
	if r := recover(); r != nil {
		util.Log.Error("Recovery error", interfaces.LogData{
			"error": r,
		})
	}
}
