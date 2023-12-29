package handler

import (
	"bwastartup-backend/helper"
	"bwastartup-backend/user"
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
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Account Account Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	u, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.ApiResponse("Account Account Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userResponse := user.ResponseUser(u, "token")

	response := helper.ApiResponse("Account has been registred", http.StatusCreated, "success", userResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	var input user.LoginUserinput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	login, err := h.userService.LoginUser(input)
	if err != nil {
		response := helper.ApiResponse("Login Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := user.ResponseUser(login, "token")
	response := helper.ApiResponse("Login Success", http.StatusOK, "success", loginResponse)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	var input user.CheckEmailInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)

		errorMessage := gin.H{"error": errors}

		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"error": "Internal Server Error"}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "failed", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{"is_available": isEmailAvailable}
	metaMessage := "Email has been registered"
	if isEmailAvailable {
		metaMessage = "Email is available"
	}
	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)

}
