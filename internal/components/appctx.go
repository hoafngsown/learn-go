package components

import (
	"learn-go/v2/internal/interfaces"
	"log"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetLog() interfaces.LogUtil
	GetLogger() *log.Logger
}

type appCtx struct {
	DB     *gorm.DB
	Log    interfaces.LogUtil
	Logger *log.Logger
}

func NewAppContext(db *gorm.DB, log interfaces.LogUtil, logger *log.Logger) *appCtx {
	return &appCtx{
		DB:     db,
		Log:    log,
		Logger: logger,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.DB
}

func (ctx *appCtx) GetLog() interfaces.LogUtil {
	return ctx.Log
}

func (ctx *appCtx) GetLogger() *log.Logger {
	return ctx.Logger
}
