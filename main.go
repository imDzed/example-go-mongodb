package main

import (
	"gomongo/config"
	uh "gomongo/features/users/handler"
	ur "gomongo/features/users/repository"
	us "gomongo/features/users/services"
	"gomongo/routes"
	"gomongo/utils/database"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	cfg := config.InitConfig()

	mongoClient, err := database.InitMongoDB(*cfg)
	if err != nil {
		e.Logger.Fatal("tidak bisa start bro", err.Error())
	}

	userRepo := ur.New(mongoClient.Client().Database(cfg.DBNAME), "contoh2")
	userService := us.New(userRepo)
	userHander := uh.New(userService)

	routes.InitRoute(e, userHander)

	e.Logger.Fatal(e.Start(":8080"))
}
