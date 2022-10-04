package db

// import (
// 	"log"

// 	registerAuthController "github.com/KadirbekSharau/mywealth-backend/controllers/auth/register"
// 	"github.com/KadirbekSharau/mywealth-backend/models"
// 	"gorm.io/gorm"
// )

// func AccountsDataMigrator(db *gorm.DB) (*models.EntityUsers) {
// 	registerRepository := registerAuthController.NewRepositoryRegister(db)
// 	registerService := registerAuthController.NewServiceRegister(registerRepository)
// 	admin := registerAuthController.InputUserRegister{
// 		Email: "admin1@gmail.com",
// 		Password: "admin532",
// 	}
// 	newAdmin, errAdmin := registerService.AdminRegisterService(&admin)
// 	if errAdmin == "REGISTER_CONFLICT_409" || errAdmin == "REGISTER_FAILED_403" {
// 		log.Println(errAdmin)
// 	}

// 	return newAdmin;
// }