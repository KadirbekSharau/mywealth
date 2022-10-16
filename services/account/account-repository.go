package accountService

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

// Get All Accounts from Repository
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

// Create Account
func (r *repository) CreateAccount(input *models.EntityAccounts) (*models.EntityAccounts, string) {
	var account models.EntityAccounts
	db := r.db.Model(&account)
	errorCode := make(chan string, 1)

	checkExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&account)

	if checkExist.RowsAffected > 0 {
		errorCode <- "CREATE_CONFLICT_409"
		return &account, <-errorCode
	}

	account.Name = input.Name
	account.Active = true
	account.UserID = input.UserID
	
	addNewAcc := r.db.Debug().Create(&account)

	

	db.Commit()

	if addNewAcc.Error != nil {
		errorCode <- "CREATE_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &account, <-errorCode
}

// Update Account Name 
func (r *repository) UpdateAccountNameByID(input *InputUpdateAccountByID) (*models.EntityAccounts, string) {
	var model models.EntityAccounts
	db := r.db.Model(&model)
	errorCode := make(chan string, 1)


	result := db.Debug().Where("id = ?", input.ID).Update("name", input.Name).Find(&model)


	if result.RowsAffected < 1 {
		errorCode <- "RESULTS_NOT_FOUND_404"
		return &model, <-errorCode
	}else if result.Error != nil {
		errorCode <- "RESULTS_NOT_FOUND_404"
		return &model, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &model, <-errorCode
}