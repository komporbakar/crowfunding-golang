package campaign

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
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
