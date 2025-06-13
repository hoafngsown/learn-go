package ginrestaurant

import (
	restaurantbiz "learn-go/v2/internal/app/restaurant/biz"
	restaurantmodel "learn-go/v2/internal/app/restaurant/model"
	restaurantstorage "learn-go/v2/internal/app/restaurant/storage"
	"learn-go/v2/internal/common"
	"learn-go/v2/internal/components"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		log := appCtx.GetUtil().Log

		log.Info("Processing update restaurant")

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid restaurant ID"})
			return
		}

		restaurant := restaurantmodel.RestaurantUpdate{}

		if err := c.ShouldBind(&restaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewUpdateRestaurantBiz(store)

		if err := biz.UpdateRestaurant(c.Request.Context(), uint(id), &restaurant); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(restaurant))
	}
}
