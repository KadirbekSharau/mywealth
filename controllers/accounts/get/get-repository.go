package getAccountsController

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

/* Get All Accounts from Repository */
func (r *repository) GetAllAccounts() (*[]models.EntityAccounts, string) {
	var accounts []models.EntityAccounts
	db := r.db.Model(&accounts)
	errorCode := make(chan string, 1)

	result := db.Debug().Select("*").Find(&accounts)

	if result.Error != nil {
		errorCode <- "RESULTS_NOT_FOUND_404"
		return &accounts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &accounts, <-errorCode
}
