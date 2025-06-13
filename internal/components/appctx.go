package components

import (
	"learn-go/v2/internal/interfaces"

	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	GetUtil() *interfaces.Util
}

type appCtx struct {
	DB   *gorm.DB
	Util *interfaces.Util
}

func NewAppContext(db *gorm.DB, util *interfaces.Util) *appCtx {
	return &appCtx{
		DB:   db,
		Util: util,
	}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.DB
}

func (ctx *appCtx) GetUtil() *interfaces.Util {
	return ctx.Util
}
