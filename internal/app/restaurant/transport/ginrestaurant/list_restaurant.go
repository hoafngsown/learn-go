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

func ListRestaurant(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMainDBConnection()
		log := appCtx.GetUtil().Log

		log.Info("Processing to get list restaurant")

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		var filter restaurantmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)

		paging.FullFill()

		result, err := biz.ListRestaurant(c.Request.Context(), &paging, &filter)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
