package main

import (
	"fmt"
	"kodelance/auth"
	"kodelance/config"
	"kodelance/handler"
	"kodelance/user"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	if err := gotenv.Load(); err != nil {
		fmt.Println(err)
		panic("Failed load env")
	}

	db, err := config.InitDb()
	if err != nil {
		fmt.Println(err)
		panic("Something wrong with database")
	}
	fmt.Println("Sukses connect ke database!")

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/auth/register", userHandler.RegisterUser)
	api.POST("/auth/login", userHandler.LoginUser)
	api.POST("/email_checkers", userHandler.IsEmailAvailable)

	router.Run()
}
