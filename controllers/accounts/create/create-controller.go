package createAccountController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	CreateAccount(input *InputCreateAccount) (*models.EntityAccounts, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) Service {
	return &service{repo: repo}
}

/* Create account controller */
func (s *service) CreateAccount(input *InputCreateAccount) (*models.EntityAccounts, string) {
	acccount := models.EntityAccounts{
		Name: input.Name,
		UserID: input.UserID,
	}

	return s.repo.CreateAccount(&acccount)
}

type InputCreateAccount struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}