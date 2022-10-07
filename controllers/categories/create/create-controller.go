package createCategoryController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	CreateCategory(input *InputCreateCategory) (*models.EntityCategories, string)
}

type service struct {
	repo *repository
}

func NewService(repo *repository) Service {
	return &service{repo: repo}
}

/* Create category controller */
func (s *service) CreateCategory(input *InputCreateCategory) (*models.EntityCategories, string) {
	category := models.EntityCategories{
		Name:   input.Name,
		UserID: input.UserID,
	}

	return s.repo.CreateCategory(&category)
}

type InputCreateCategory struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}
