package handler

import (
	"bwastartup-backend/campaign"
	"bwastartup-backend/helper"
	"bwastartup-backend/user"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))
	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {

		response := helper.ApiResponse("Campaign to get campaigns", http.StatusBadRequest, "error", nil)

		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to get campaigns", http.StatusOK, "success", campaign.ResponseCampaigns(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.GetCampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignDetail, err := h.service.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get detail campaign", http.StatusNotFound, "error", nil)
		c.JSON(http.StatusNotFound, response)
		return
	}

	response := helper.ApiResponse("Success to get detail campaign", http.StatusOK, "success", campaign.ResponseDetailCampaign(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) CreateCampaign(ctx *gin.Context) {
	var input campaign.CreateCampaignInput
	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		fmt.Println(errors)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to create campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.service.CreateCampaign(input)
	if err != nil {
		response := helper.ApiResponse("Failed to create campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Success to create campaign", http.StatusOK, "success", campaign.ResponseCampaign(newCampaign))
	ctx.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UpdateCampaign(ctx *gin.Context) {
	var inputId campaign.GetCampaignDetailInput

	err := ctx.ShouldBindUri(&inputId)
	if err != nil {
		response := helper.ApiResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var inputData campaign.CreateCampaignInput
	err = ctx.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		fmt.Println(errors)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Failed to update campaign", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedCampaign, err := h.service.UpdateCampaign(inputId, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to update campaign", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to update campaign", http.StatusOK, "success", campaign.ResponseCampaign(updatedCampaign))
	ctx.JSON(http.StatusOK, response)
}
