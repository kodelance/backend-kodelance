package routes

import (
	"kodelance/auth"
	"kodelance/handler"
	"kodelance/helper"
	"kodelance/user"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type routes struct {
	userHandler handler.UserHandler
	userSevice  user.Service
	authService auth.Service
}

func NewRoutes(userHandler handler.UserHandler, userSevice user.Service, authService auth.Service) *routes {
	return &routes{userHandler, userSevice, authService}
}

func (r *routes) Route() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	api := router.Group("/api/v1")
	{
		api.GET("/", func(c *gin.Context) {
			response := helper.ApiResponse("Selamat datang di Kodelance!", 200, "success", nil)
			c.JSON(200, response)
		})
		api.POST("/auth/register", r.userHandler.RegisterUser)
		api.POST("/auth/login", r.userHandler.LoginUser)
		api.POST("/email_checkers", r.userHandler.IsEmailAvailable)
		api.POST("/auth/test", authMiddleware(r.authService, r.userSevice), r.userHandler.TestAuth)
	}
	return router
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		arrToken := strings.Split(authHeader, " ")
		if len(arrToken) != 2 {
			response := helper.ApiResponse("Wrong Token", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		tokenString := arrToken[1]

		token, err := authService.ValidationToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Token Undifined", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.ApiResponse("Token Undifined", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := uint(claim["jwt_user_id"].(float64))

		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.ApiResponse("User not found by Id", http.StatusNotFound, "error", nil)
			c.AbortWithStatusJSON(http.StatusNotFound, response)
			return
		}

		c.Set("userLoggedIn", user)
	}
}
