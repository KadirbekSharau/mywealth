package createAccountController

import (
	"github.com/KadirbekSharau/mywealth-backend/models"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewCreateAccountRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateAccountRepository(input *models.EntityAccounts) (*models.EntityAccounts, string) {
	var account models.EntityAccounts
	db := r.db.Model(&account)
	errorCode := make(chan string, 1)

	checkExist := db.Debug().Select("*").Where("name = ?", input.Name).Find(&account)

	if checkExist.RowsAffected > 0 {
		errorCode <- "CREATE_ACCOUNT_CONFLICT_409"
		return &account, <-errorCode
	}

	account.Name = input.Name
	account.Active = true
	account.UserID = input.UserID
	
	addNewAcc := r.db.Debug().Create(&account)

	

	db.Commit()

	if addNewAcc.Error != nil {
		errorCode <- "CREATE_ACCOUNT_FAILED_403"
	} else {
		errorCode <- "nil"
	}

	return &account, <-errorCode
}