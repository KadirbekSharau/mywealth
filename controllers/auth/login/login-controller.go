package loginAuthController

import (
	model "github.com/KadirbekSharau/mywealth-backend/models"
)

type Service interface {
	UserLoginService(input *InputLogin) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceLogin(repository Repository) *service {
	return &service{repository: repository}
}

/* User Login Service */
func (s *service) UserLoginService(input *InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repository.UserLoginRepository(&user)

	return resultLogin, errLogin
}