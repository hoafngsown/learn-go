package ginupload

import (
	"fmt"
	"learn-go/v2/internal/common"
	components "learn-go/v2/internal/components/app_context"
	"net/http"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func CheckError(err error, errFnc func(err error) *common.AppError) {
	if err != nil {
		panic(errFnc(err))
	}
}

func UploadImageStatic(appCtx components.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		CheckError(err, common.ErrorInvalidRequest)

		fileExt := filepath.Ext(fileHeader.Filename)
		fileName := fmt.Sprintf("%s/%d%s", "static", time.Now().Nanosecond(), fileExt)

		err = c.SaveUploadedFile(fileHeader, fileName)
		CheckError(err, common.ErrorCannotSaveFile)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse("Upload success"))
	}
}
