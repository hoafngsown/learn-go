package main

import (
	"fmt"
	"learn-go/v2/internal/app/restaurant/transport/ginrestaurant"
	"learn-go/v2/internal/components"
	"learn-go/v2/internal/interfaces"
	"learn-go/v2/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	Name       string
	Address    string
	Phone      string
	Email      string
	Rating     float64
	Cuisine    string
	PriceRange string
}

func main() {
	env := NewEnvironment()
	dbConnStr := env.GetEnv(interfaces.DATABASE, "CONN_STR")

	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})

	if err != nil {
		fmt.Printf("Connect to database failed: %v\n", err)
		panic(err)
	}

	db.Debug()

	log := NewLogger()
	util := NewUtil(log, env)

	appCtx := components.NewAppContext(db, util)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	{
		v1 := router.Group("/v1")
		v1.GET("/restaurants", ginrestaurant.ListRestaurant(appCtx))
		v1.GET("/restaurants/:id", ginrestaurant.DetailRestaurant(appCtx))
		v1.POST("/restaurants", ginrestaurant.CreateRestaurant(appCtx))
		v1.PATCH("/restaurants/:id", ginrestaurant.UpdateRestaurant(appCtx))
		v1.DELETE("/restaurants/:id", ginrestaurant.DeleteRestaurant(appCtx))
	}

	serverPort := env.GetEnv(interfaces.SERVER, "PORT")
	router.Run(fmt.Sprintf(":%s", serverPort))
}

func NewUtil(log *logging, env *environment) *interfaces.Util {
	return &interfaces.Util{
		Log:         log,
		Logger:      log.ErrorLogger,
		Environment: env,
	}
}
