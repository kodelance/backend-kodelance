package main

import (
	"fmt"
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
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/auth/login", userHandler.LoginUser)

	router.Run()
}
