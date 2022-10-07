package getAccountsController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	GetAllAccounts() (*[]models.EntityAccounts, string)
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