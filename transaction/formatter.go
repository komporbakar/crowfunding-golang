package transaction

import (
	"time"
)

type CampaignTransactionResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func ResponseCampaignTransactions(transaction Transaction) CampaignTransactionResponse {
	formatter := CampaignTransactionResponse{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionResponse {
	if len(transactions) == 0 {
		return []CampaignTransactionResponse{}
	}

	var transactionFormatter []CampaignTransactionResponse

	for _, transaction := range transactions {
		formatter := ResponseCampaignTransactions(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}
	return transactionFormatter
}

type UserTransactionResponse struct {
	Id        int              `json:"id"`
	Amount    int              `json:"amount"`
	Status    string           `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	Campaign  CampaignResponse `json:"campaign"`
}

type CampaignResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}

func ResponseUserTransactions(transaction Transaction) UserTransactionResponse {
	formatter := UserTransactionResponse{}
	formatter.Id = transaction.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignResponse{}
	campaignFormatter.Name = transaction.Campaign.Name
	campaignFormatter.ImageUrl = ""
	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName
	}

	formatter.Campaign = campaignFormatter

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionResponse {
	if len(transactions) == 0 {
		return []UserTransactionResponse{}
	}

	var transactionFormatter []UserTransactionResponse

	for _, transaction := range transactions {
		formatter := ResponseUserTransactions(transaction)
		transactionFormatter = append(transactionFormatter, formatter)
	}
	return transactionFormatter
}
