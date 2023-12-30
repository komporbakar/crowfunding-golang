package handler

import (
	"bwastartup-backend/auth"
	"bwastartup-backend/helper"
	"bwastartup-backend/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
	authService auth.Service
}

func NewUserHandler(userService user.Service, authService auth.Service) *userHandler {
	return &userHandler{userService, authService}
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

	token, err := h.authService.GenerateToken(u.Id)
	if err != nil {
		response := helper.ApiResponse("Account Account Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userResponse := user.ResponseUser(u, token)

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
	loginUser, err := h.userService.LoginUser(input)
	if err != nil {
		response := helper.ApiResponse("Login Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	token, err := h.authService.GenerateToken(loginUser.Id)
	if err != nil {
		response := helper.ApiResponse("Login Failed", http.StatusBadRequest, "failed", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	loginResponse := user.ResponseUser(loginUser, token)
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

func (h *userHandler) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")

	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	userId := 1

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)
	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userId, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Failed to upload avatar", http.StatusBadRequest, "failed", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Avatar Successfully Uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
