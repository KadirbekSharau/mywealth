package createCategoryController

import "github.com/KadirbekSharau/mywealth-backend/models"

type Service interface {
	CreateCategoryService(input *InputCreateCategory) (*models.EntityCategories, string)
}

type service struct {
	repository repository
}

func NewCreateCategoryService(repository repository) Service {
	return &service{repository: repository}
}

/* Create category controller */
func (s *service) CreateCategoryService(input *InputCreateCategory) (*models.EntityCategories, string) {
	category := models.EntityCategories{
		Name:   input.Name,
		UserID: input.UserID,
	}

	return s.repository.CreateCategoryRepository(&category)
}

type InputCreateCategory struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}
