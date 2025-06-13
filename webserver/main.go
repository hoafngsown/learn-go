package main

import (
	"fmt"
	"learn-go/v2/module/component"
	"learn-go/v2/module/interfaces"
	"net/http"

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
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	env := NewEnvironment()
	dbConnStr := env.GetEnv(interfaces.DATABASE, "CONN_STR")

	db, err := gorm.Open(mysql.Open(dbConnStr), &gorm.Config{})

	if err != nil {
		fmt.Printf("Connect to database failed: %v\n", err)
		panic(err)
	}

	log := NewLogger()
	util := NewUtil(log, env)

	appCtx := component.NewAppContext(db, util)

	fmt.Println("Connect to database successfully", appCtx)

	serverPort := env.GetEnv(interfaces.SERVER, "PORT")

	r.Run(fmt.Sprintf(":%s", serverPort))
}

func NewUtil(log *Logger, env *Environment) *interfaces.Util {
	return &interfaces.Util{
		Log:         log,
		Logger:      log.ErrorLogger,
		Environment: env,
	}
}
