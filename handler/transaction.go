package handler

import (
	"bwastartup-backend/helper"
	"bwastartup-backend/transaction"
	"bwastartup-backend/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransaction(ctx *gin.Context) {
	var input transaction.GetCampaignTransactionInput

	err := ctx.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return

	}

	currentUser := ctx.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to get detail campaign", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	ctx.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(user.User)
	userId := currentUser.Id

	transactions, err := h.service.GetTransactionByUserId(userId)
	if err != nil {
		response := helper.ApiResponse("Failed to get user's transactions", http.StatusBadRequest, "error", nil)
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Success to get user's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	ctx.JSON(http.StatusOK, response)
}
