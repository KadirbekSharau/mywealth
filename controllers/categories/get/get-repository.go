package getCategoriesController

import (
	"github.com/KadirbekSharau/mywealth-backend/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Get All Categories from Repository */
func (r *repository) GetAllCategories() (*[]models.EntityCategories, string) {
	var categories []models.EntityCategories
	db := r.db.Model(&categories)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Find(&categories)

	if result.Error != nil {
		errorCode <- "RESULTS_NOT_FOUND_404"
		return &categories, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &categories, <-errorCode
}
