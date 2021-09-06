package handler

import (
	"kodelance/helper"
	"kodelance/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var userInput user.UserInput

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

	formatData := user.FormatterOutput(newUser, "tokentoken")
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

	formatData := user.FormatterOutput(userLogged, "tokentokentoken")
	response := helper.ApiResponse("Successfuly Loggedin", 200, "success", formatData)
	c.JSON(http.StatusOK, response)
}
