package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/controllers/categories/create"
	"github.com/KadirbekSharau/mywealth-backend/handlers/categories/create"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Categories routes */
func InitCategoryRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createRepository  = createCategoryController.NewCreateCategoryRepository(db)
		createService = createCategoryController.NewCreateCategoryService(*createRepository)
		createHandler = createCategoryHandler.NewCreateCategoryHandler(createService)
	)

	groupRoute := route.Group("/api/v1/categories")
	groupRoute.POST("/create", createHandler.CreateCategoryHandler)
}