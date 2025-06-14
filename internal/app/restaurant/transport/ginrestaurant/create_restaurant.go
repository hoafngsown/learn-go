package ginrestaurant

import (
	restaurantbiz "learn-go/v2/internal/app/restaurant/biz"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	restaurantstorage "learn-go/v2/internal/app/restaurant/storage"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/components"

	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		log := appCtx.GetLog()

		log.Info("Processing to create restaurant")

		restaurant := restaurantmodel.RestaurantCreate{}

		if err := c.ShouldBindJSON(&restaurant); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateRestaurant(c.Request.Context(), &restaurant); err != nil {
			panic(err)
		}

		restaurant.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurant.FakeId))
	}
}
