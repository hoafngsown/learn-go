package ginrestaurant

import (
	restaurantbiz "learn-go/v2/internal/app/restaurant/biz"
	restaurantstorage "learn-go/v2/internal/app/restaurant/storage"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/components"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		log := appCtx.GetUtil().Log

		log.Info("Processing delete restaurant")

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewDeleteRestaurantBiz(store)

		if err := biz.DeleteRestaurant(c.Request.Context(), uint(id)); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("ok"))
	}
}
