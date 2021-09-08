package main

import (
	"fmt"
	"kodelance/auth"
	"kodelance/config"
	"kodelance/handler"
	"kodelance/routes"
	"kodelance/user"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// Setup Env File
	var APP_ENV = os.Getenv("APP_ENV")
	if APP_ENV != "PRODUCTION" {
		if err := godotenv.Load(); err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Setup Database
	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		panic("Something wrong with database")
	}
	fmt.Println("Sukses connect ke database!")

	// Setup Repository
	userRepository := user.NewRepository(db)

	// Setup Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	// Setup Handler
	userHandler := handler.NewUserHandler(userService, authService)

	// Setup Router
	r := routes.NewRoutes(userHandler, userService, authService)
	route := r.Route()

	route.Run()
}
