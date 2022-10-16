package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/services/account"
	"github.com/KadirbekSharau/mywealth-backend/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Accounts routes */
func InitAccountsRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository  = accountService.NewRepository(db)
		service = accountService.NewService(repository)
		handler = handlers.NewAccountHandler(service)
	)

	groupRoute := route.Group("/api/v1/accounts")
	groupRoute.POST("/create", handler.CreateAccount)
	groupRoute.GET("/get", handler.GetAllAcounts)
}