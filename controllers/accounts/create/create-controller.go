package createAccountController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	CreateAccountService(input *InputCreateAccount) (*models.EntityAccounts, string)
}

type service struct {
	repository repository
}

func NewCreateAccountService(repository repository) Service {
	return &service{repository: repository}
}

/* Create account controller */
func (s *service) CreateAccountService(input *InputCreateAccount) (*models.EntityAccounts, string) {
	acccount := models.EntityAccounts{
		Name: input.Name,
		UserID: input.UserID,
	}

	return s.repository.CreateAccountRepository(&acccount)
}

type InputCreateAccount struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}