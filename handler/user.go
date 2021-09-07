package handler

import (
	"kodelance/auth"
	"kodelance/helper"
	"kodelance/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	RegisterUser(c *gin.Context)
	LoginUser(c *gin.Context)
	IsEmailAvailable(c *gin.Context)
	TestAuth(c *gin.Context)
}

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
}

func (h *userHandler) TestAuth(c *gin.Context) {
	userLogged := c.MustGet("userLoggedIn").(user.User)
	response := helper.ApiResponse("Halo user", 200, "success", user.FormatterOutput(userLogged, "haloyo"))
	c.JSON(http.StatusOK, response)
	return
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var userInput user.RegisterInput

	if err := c.ShouldBindJSON(&userInput); err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Register Failed", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	newUser, err := h.userService.RegisterUser(userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Register Failed", 400, "error", nil))
		return
	}

	token, err := h.authService.GenerateToken(newUser.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Register Failed", 400, "error", nil))
		return
	}

	formatData := user.FormatterOutput(newUser, token)
	response := helper.ApiResponse("Account has been registed", 201, "success", formatData)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var loginInput user.LoginInput

	if err := c.ShouldBindJSON(&loginInput); err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	userLogged, err := h.userService.LoginUser(loginInput)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Login Failed", 400, "error", errorMessage))
		return
	}

	token, err := h.authService.GenerateToken(userLogged.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Register Failed", 400, "error", nil))
		return
	}

	formatData := user.FormatterOutput(userLogged, token)
	response := helper.ApiResponse("Successfuly Loggedin", 200, "success", formatData)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) IsEmailAvailable(c *gin.Context) {
	var userEmail user.CheckEmailInput

	if err := c.ShouldBindJSON(&userEmail); err != nil {
		errors := helper.FormatError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	isAvailable, err := h.userService.IsEmailAvailable(userEmail)
	if err != nil {
		errorMessage := gin.H{"errors": "Server Error"}
		c.JSON(http.StatusBadRequest, helper.ApiResponse("Email Checking Failed", http.StatusUnprocessableEntity, "error", errorMessage))
		return
	}

	metaMessage := "Email Tidak Tersedia"
	if isAvailable {
		metaMessage = "Email Tersedia"
	}

	data := gin.H{
		"is_available": isAvailable,
	}

	c.JSON(http.StatusOK, helper.ApiResponse(metaMessage, http.StatusOK, "success", data))
}
