package main

import (
	"fmt"
	"learn-go/v2/internal/app/restaurant/transport/ginrestaurant"
	"learn-go/v2/internal/app/upload/transport/ginupload"
	components "learn-go/v2/internal/components/app_context"
	"learn-go/v2/internal/components/environment"
	"learn-go/v2/internal/components/logger"
	uploadprovider "learn-go/v2/internal/components/upload_provider"
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
	env := environment.NewEnvironment()
	dbConnStr := env.GetEnv(environment.DATABASE, "CONN_STR")

	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})

	if err != nil {
		fmt.Printf("Connect to database failed: %v\n", err)
		panic(err)
	}

	db.Debug()

	log := logger.NewLogger()
	provider := uploadprovider.NewS3Provider(
		env.GetEnv(environment.UPLOAD_PROVIDER, "BUCKET_NAME"),
		env.GetEnv(environment.UPLOAD_PROVIDER, "REGION"),
		env.GetEnv(environment.UPLOAD_PROVIDER, "API_KEY"),
		env.GetEnv(environment.UPLOAD_PROVIDER, "SECRET_KEY"),
		env.GetEnv(environment.UPLOAD_PROVIDER, "DOMAIN"),
	)

	appCtx := components.NewAppContext(db, log, provider)

	router := gin.Default()
	router.Use(middleware.Recover(appCtx))
	router.Static("/static", "./static")

	{
		v1 := router.Group("/v1")

		{

			v1.GET("/restaurants", ginrestaurant.ListRestaurant(appCtx))
			v1.GET("/restaurants/:id", ginrestaurant.DetailRestaurant(appCtx))
			v1.POST("/restaurants", ginrestaurant.CreateRestaurant(appCtx))
			v1.PATCH("/restaurants/:id", ginrestaurant.UpdateRestaurant(appCtx))
			v1.DELETE("/restaurants/:id", ginrestaurant.DeleteRestaurant(appCtx))
		}

		{
			v1.POST("/upload", ginupload.UploadImage(appCtx))
			v1.POST("/upload/static", ginupload.UploadImageStatic(appCtx))
		}
	}

	serverPort := env.GetEnv(environment.SERVER, "PORT")
	router.Run(fmt.Sprintf(":%s", serverPort))
}
