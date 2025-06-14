package ginrestaurant

import (
	restaurantbiz "learn-go/v2/internal/app/restaurant/biz"
	restaurantstorage "learn-go/v2/internal/app/restaurant/storage"
	"learn-go/v2/internal/common"
	components "learn-go/v2/internal/components/app_context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DetailRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		log := appCtx.GetLog()

		log.Info("Processing detail restaurant")

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDetailRestaurantBiz(store)

		restaurant, err := biz.DetailRestaurant(c.Request.Context(), uint(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		restaurant.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurant))
	}
}
