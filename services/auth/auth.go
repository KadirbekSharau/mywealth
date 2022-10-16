package authService

import (
	model "github.com/KadirbekSharau/mywealth-backend/models"
)

type Service interface {
	UserLogin(input *InputLogin) (*model.EntityUsers, string)
	ActiveUserRegister(input *InputUserRegister) (*model.EntityUsers, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* User Login Service */
func (s *service) UserLogin(input *InputLogin) (*model.EntityUsers, string) {

	user := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}

	resultLogin, errLogin := s.repo.UserLogin(&user)

	return resultLogin, errLogin
}

/* Active User Registration Service */
func (s *service) ActiveUserRegister(input *InputUserRegister) (*model.EntityUsers, string) {

	users := model.EntityUsers{
		Email:    input.Email,
		Password: input.Password,
	}
	resultRegister, errRegister := s.repo.ActiveUserRegisterRepository(&users)

	return resultRegister, errRegister
}