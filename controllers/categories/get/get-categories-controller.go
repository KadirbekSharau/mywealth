package getCategoriesController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	GetAllCategories() (*[]models.EntityCategories, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) *service {
	return &service{repo: repo}
}

/* Get All Categories */
func (s *service) GetAllCategories() (*[]models.EntityCategories, string) {

	res, err := s.repo.GetAllCategories()

	return res, err
}
