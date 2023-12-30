package campaign

type CampaignResponse struct {
	Id               int    `json:"id"`
	UserId           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func ResponseCampaign(campaign Campaign) CampaignResponse {
	campaignResponse := CampaignResponse{}
	campaignResponse.Id = campaign.Id
	campaignResponse.UserId = campaign.UserId
	campaignResponse.Name = campaign.Name
	campaignResponse.ShortDescription = campaign.ShortDescription
	campaignResponse.GoalAmount = campaign.GoalAmount
	campaignResponse.CurrentAmount = campaign.CurrentAmount
	campaignResponse.Slug = campaign.Slug

	campaignResponse.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignResponse.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignResponse
}

func ResponseCampaigns(campaigns []Campaign) []CampaignResponse {

	campaignsResponse := []CampaignResponse{}
	for _, campaign := range campaigns {
		campaignResponse := ResponseCampaign(campaign)
		campaignsResponse = append(campaignsResponse, campaignResponse)
	}

	return campaignsResponse
}
