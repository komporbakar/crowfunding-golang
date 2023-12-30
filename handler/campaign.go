package handler

import (
	"bwastartup-backend/campaign"
	"bwastartup-backend/helper"
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
