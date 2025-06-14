package components

import (
	"learn-go/v2/internal/components/logger"
	uploadprovider "learn-go/v2/internal/components/upload_provider"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetLog() logger.LogUtil
	GetUploadProvider() uploadprovider.UploadProvider
}

type appCtx struct {
	DB             *gorm.DB
	Log            logger.LogUtil
	UploadProvider uploadprovider.UploadProvider
}

func NewAppContext(db *gorm.DB, log logger.LogUtil, uploadProvider uploadprovider.UploadProvider) *appCtx {
	return &appCtx{
		DB:             db,
		Log:            log,
		UploadProvider: uploadProvider,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.DB
}

func (ctx *appCtx) GetLog() logger.LogUtil {
	return ctx.Log
}

func (ctx *appCtx) GetUploadProvider() uploadprovider.UploadProvider {
	return ctx.UploadProvider
}
