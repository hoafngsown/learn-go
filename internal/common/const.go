package common

import "learn-go/v2/internal/components/logger"

const DBTypeRestaurant = 1

func AppRecover(log logger.LogUtil) {
	if r := recover(); r != nil {
		log.Error("Recovery error", logger.LogData{
			"error": r,
		})
	}
}
