package main

import (
	"auth-service/api"
	"auth-service/api/handler"
	"auth-service/pkg/config"
	"auth-service/pkg/logs"
	"auth-service/service"
	"auth-service/storage/postgres"
	"log"
)

func main() {
	logger := logs.InitLogger()
	cofg := config.Load()

	db, err := postgres.ConnectPostgres(cofg)
	if err != nil {
		logger.Error("Error connecting to database", "error", err)
		log.Fatal(err)
	}

	authSt := postgres.NewAuthRepo(db)

	authSr := service.NewAuthService(authSt, logger)

	authHd := handler.NewAuthHandler(logger, authSr)

	router := api.NewRouter(authHd)

	router.InitRouter()
	err = router.RunRouter(cofg)

	if err != nil {
		logger.Error("Error starting server", "error", err)
		log.Fatal(err)
	}
}
