package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	LoginUser(input LoginUserinput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
	GetUserById(Id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *service) LoginUser(input LoginUserinput) (User, error) {
	email := input.Email
	password := input.Password
	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("no user found on that email address")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}
	return user, nil
}

func (service *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	u, err := service.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}
	if u.Id == 0 {
		return true, nil
	}
	return false, nil
}

func (s *service) SaveAvatar(Id int, fileLocation string) (User, error) {
	user, err := s.repository.FindById(Id)
	if err != nil {
		return user, err
	}

	user.AvatarFileName = fileLocation

	u, err := s.repository.Update(user)
	if err != nil {
		return u, err
	}

	return u, nil
}

func (s *service) GetUserById(Id int) (User, error) {
	user, err := s.repository.FindById(Id)
	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}
