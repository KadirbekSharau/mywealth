package registerAuthController

import (
	model "github.com/KadirbekSharau/mywealth-backend/models"
	"gorm.io/gorm"
)

type Repository interface {
	ActiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryRegister(db *gorm.DB) *repository {
	return &repository{db: db}
}

/* Active User Registration Repository */
func (r *repository) ActiveUserRegisterRepository(input *model.EntityUsers) (*model.EntityUsers, string) {

	var users model.EntityUsers
	db := r.db.Model(&users)
	errorCode := make(chan string, 1)

	checkUserAccount := db.Debug().Select("*").Where("email = ?", input.Email).Find(&users)

	if checkUserAccount.RowsAffected > 0 {
		errorCode <- "REGISTER_CONFLICT_409"
		return &users, <-errorCode
	}

	users.Name = input.Name
	users.Surname = input.Surname
	users.Email = input.Email
	users.Password = input.Password
	users.Active = true

	addNewUser := db.Debug().Create(&users)
	db.Commit()

	if addNewUser.Error != nil {
		errorCode <- "REGISTER_FAILED_403"
		return &users, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &users, <-errorCode
}