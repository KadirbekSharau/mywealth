package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/controllers/accounts/create"
	"github.com/KadirbekSharau/mywealth-backend/handlers/accounts/create"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Accounts routes */
func InitAccountsRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createRepository  = createAccountController.NewCreateAccountRepository(db)
		createService = createAccountController.NewCreateAccountService(*createRepository)
		createHandler = createAccountHandler.NewCreateAccountHandler(createService)
	)

	groupRoute := route.Group("/api/v1/accounts")
	groupRoute.POST("/create", createHandler.CreateAccountHandler)
}