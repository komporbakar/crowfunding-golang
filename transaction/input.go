package transaction

import "bwastartup-backend/user"

type GetCampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
