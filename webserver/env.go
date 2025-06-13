package main

import (
	"fmt"
	"learn-go/v2/module/interfaces"
	"os"

	"github.com/joho/godotenv"
)

var _ interfaces.Environment = (*Environment)(nil)

const (
	APP_DEVELOPMENT = "development"
	APP_PRODUCTION  = "production"
)

type Environment struct{}

func NewEnvironment() *Environment {
	appEnv := os.Getenv("APP_ENV")
	fmt.Printf("APP_ENV: %s\n", appEnv)

	if appEnv == APP_DEVELOPMENT {
		loadAppDev()
	} else {
		loadAppProd()
	}

	return &Environment{}
}

func (e *Environment) GetEnv(space interfaces.SPACE, key string) string {
	spacePrefix := map[interfaces.SPACE]interfaces.SPACE{
		interfaces.SERVER:   "SERVER_",
		interfaces.DATABASE: "DB_",
	}

	prefix, ok := spacePrefix[space]

	if !ok {
		panic(fmt.Sprintf("Invalid space: %s", space))
	}

	environmentKey := fmt.Sprintf("%s%s", prefix, key)
	value := os.Getenv(environmentKey)

	if value == "" {
		panic(fmt.Sprintf("Environment variable %s not set", environmentKey))
	}

	return value
}

func loadAppDev() {
	if err := godotenv.Load(".env.development"); err != nil {
		if err := godotenv.Load("../.env.development"); err != nil {
			panic(fmt.Sprintf("Error loading .env.development file: %v", err))
		}
	}
}

func loadAppProd() {
	godotenv.Load()
}
