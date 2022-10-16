package categoryService

import (
	"github.com/KadirbekSharau/mywealth-backend/models"
)

type Service interface {
	CreateCategory(input *InputCreateCategory) (*models.EntityCategories, string)
	GetAllCategories() (*[]models.EntityCategories, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) Service {
	return &service{repo: repo}
}

/* Get All Categories */
func (s *service) GetAllCategories() (*[]models.EntityCategories, string) {

	res, err := s.repo.GetAllCategories()

	return res, err
}

/* Create category controller */
func (s *service) CreateCategory(input *InputCreateCategory) (*models.EntityCategories, string) {
	category := models.EntityCategories{
		Name:   input.Name,
		UserID: input.UserID,
	}

	return s.repo.CreateCategory(&category)
}