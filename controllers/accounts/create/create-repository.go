package createAccountController

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