package createCategoryController

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

func (r *repository) CreateCategory(input *models.EntityCategories) (*models.EntityCategories, string) {
	var category models.EntityCategories
	db := r.db.Model(&category)
	errorCode := make(chan string, 1)

	checkExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&category)

	if checkExist.RowsAffected > 0 {
		errorCode <- "CREATE_CONFLICT_409"
		return &category, <-errorCode
	}

	category.Name = input.Name
	category.UserID = input.UserID
	
	addNewAcc := r.db.Debug().Create(&category)

	

	db.Commit()

	if addNewAcc.Error != nil {
		errorCode <- "CREATE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &category, <-errorCode
}