package component

import (
	"learn-go/v2/module/interfaces"

	"gorm.io/gorm"
)

type appContext struct {
	DB   *gorm.DB
	Util *interfaces.Util
}

func NewAppContext(db *gorm.DB, util *interfaces.Util) *appContext {
	return &appContext{
		DB:   db,
		Util: util,
	}
}

func (ctx *appContext) GetMainDBConnection() *gorm.DB {
	return ctx.DB
}

func (ctx *appContext) GetUtil() *interfaces.Util {
	return ctx.Util
}
