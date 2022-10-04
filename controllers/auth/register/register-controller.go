package registerAuthController

import (
	model "github.com/KadirbekSharau/mywealth-backend/models"
)

type Service interface {
	ActiveUserRegisterService(input *InputUserRegister) (*model.EntityUsers, string)
}

type service struct {
	repository Repository
}

func NewServiceRegister(repository Repository) *service {
	return &service{repository: repository}
}

/* Active User Registration Service */
func (s *service) ActiveUserRegisterService(input *InputUserRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}
	resultRegister, errRegister := s.repository.ActiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}