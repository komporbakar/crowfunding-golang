package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) GetCampaigns(UserId int) ([]Campaign, error) {
	if UserId != 0 {
		campaigns, err := service.repository.FindByUserId(UserId)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := service.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil

}

func (service *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := service.repository.FindById(input.Id)
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserId = input.User.Id

	// slug created
	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.Id)
	slug := slug.Make(slugCandidate)
	campaign.Slug = slug

	newCampaign, err := s.repository.CreateCampaign(campaign)
	if err != nil {
		return campaign, err
	}
	return newCampaign, nil
}

func (s *service) UpdateCampaign(inputId GetCampaignDetailInput, inputData CreateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(inputId.Id)
	if err != nil {
		return campaign, err
	}

	if campaign.UserId != inputData.User.Id {
		return campaign, errors.New("not an owner of the campaign")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.Perks = inputData.Perks
	campaign.GoalAmount = inputData.GoalAmount

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}
	return updatedCampaign, nil
}
