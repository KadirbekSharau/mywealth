package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/controllers/accounts/create"
	"github.com/KadirbekSharau/mywealth-backend/handlers/accounts/create"
	"github.com/KadirbekSharau/mywealth-backend/controllers/accounts/get"
	"github.com/KadirbekSharau/mywealth-backend/handlers/accounts/get"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Accounts routes */
func InitAccountsRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createRepository  = createAccountController.NewRepository(db)
		createService = createAccountController.NewService(createRepository)
		createHandler = createAccountHandler.NewHandler(createService)

		getRepository  = getAccountsController.NewRepository(db)
		getService = getAccountsController.NewService(getRepository)
		getHandler = getAccountsHandler.NewHandler(getService)
	)

	groupRoute := route.Group("/api/v1/accounts")
	groupRoute.POST("/create", createHandler.CreateAccount)
	groupRoute.GET("/get", getHandler.GetAllAcounts)
}