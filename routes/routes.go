package routes

import (
	"kodelance/handler"

	"github.com/gin-gonic/gin"
)

type routes struct {
	userHandler handler.UserHandler
}

func NewRoutes(userHandler handler.UserHandler) *routes {
	return &routes{userHandler}
}

func (r *routes) Route() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.POST("/auth/register", r.userHandler.RegisterUser)
		api.POST("/auth/login", r.userHandler.LoginUser)
		api.POST("/email_checkers", r.userHandler.IsEmailAvailable)
	}
	return router
}
