package accountService

import (
	"github.com/KadirbekSharau/mywealth-backend/models"
)

type Service interface {
	GetAllAccounts() (*[]models.EntityAccounts, string)
	CreateAccount(input *InputCreateAccount) (*models.EntityAccounts, string)
	UpdateAccountNameByID(input *InputUpdateAccountByID) (*models.EntityAccounts, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* Get All Accounts */
func (s *service) GetAllAccounts() (*[]models.EntityAccounts, string) {

	res, err := s.repo.GetAllAccounts()

	return res, err
}

/* Create account controller */
func (s *service) CreateAccount(input *InputCreateAccount) (*models.EntityAccounts, string) {
	acccount := models.EntityAccounts{
		Name: input.Name,
		UserID: input.UserID,
	}

	return s.repo.CreateAccount(&acccount)
}

/* Update Account Name by ID */
func (s *service) UpdateAccountNameByID(input *InputUpdateAccountByID) (*models.EntityAccounts, string) {

	result, err := s.repo.UpdateAccountNameByID(input)

	return result, err
}