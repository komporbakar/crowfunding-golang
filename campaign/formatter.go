package campaign

import "strings"

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

type DetailCampaignResponse struct {
	Id               int                     `json:"id"`
	Name             string                  `json:"name"`
	ShortDescription string                  `json:"short_description"`
	Desciption       string                  `json:"desciption"`
	ImageUrl         string                  `json:"image_url"`
	GoalAmount       int                     `json:"goal_amount"`
	CurrentAmount    int                     `json:"current_amount"`
	UserId           int                     `json:"user_id"`
	Slug             string                  `json:"slug"`
	Perks            []string                `json:"perks"`
	User             CampaignUserResponse    `json:"user"`
	Images           []CampaignImageResponse `json:"images"`
}

type CampaignUserResponse struct {
	Name     string `json:"name"`
	ImageUrl string `json:"image_url"`
}
type CampaignImageResponse struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func ResponseDetailCampaign(campaign Campaign) DetailCampaignResponse {
	campaignResponse := DetailCampaignResponse{}
	campaignResponse.Id = campaign.Id
	campaignResponse.Name = campaign.Name
	campaignResponse.UserId = campaign.UserId
	campaignResponse.ShortDescription = campaign.ShortDescription
	campaignResponse.Desciption = campaign.Description
	campaignResponse.GoalAmount = campaign.GoalAmount
	campaignResponse.CurrentAmount = campaign.CurrentAmount
	campaignResponse.Slug = campaign.Slug
	campaignResponse.ImageUrl = ""

	if len(campaign.CampaignImages) > 0 {
		campaignResponse.ImageUrl = campaign.CampaignImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignResponse.Perks = perks

	user := campaign.User
	campaignUserFormatter := CampaignUserResponse{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageUrl = user.AvatarFileName

	campaignResponse.User = campaignUserFormatter

	images := []CampaignImageResponse{}

	for _, image := range campaign.CampaignImages {
		campaignImageResponse := CampaignImageResponse{}
		campaignImageResponse.ImageUrl = image.FileName

		isPrimary := false

		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImageResponse.IsPrimary = isPrimary

		images = append(images, campaignImageResponse)
	}

	campaignResponse.Images = images

	return campaignResponse
}
