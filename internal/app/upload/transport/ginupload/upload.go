package ginupload

import (
	"io"
	uploadbiz "learn-go/v2/internal/app/upload/biz"
	"learn-go/v2/internal/common"
	components "learn-go/v2/internal/components/app_context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadImage(ctx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		defer file.Close()

		dataBytes, err := io.ReadAll(file)
		if err != nil {
			panic(common.ErrorInvalidRequest(err))
		}

		biz := uploadbiz.NewUploadBiz(ctx.GetUploadProvider())
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
