package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/controllers/categories/create"
	"github.com/KadirbekSharau/mywealth-backend/handlers/categories/create"

	"github.com/KadirbekSharau/mywealth-backend/controllers/categories/get"
	"github.com/KadirbekSharau/mywealth-backend/handlers/categories/get"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Categories routes */
func InitCategoryRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		createRepository = createCategoryController.NewRepository(db)
		createService    = createCategoryController.NewService(createRepository)
		createHandler    = createCategoryHandler.NewHandler(createService)

		getRepository = getCategoriesController.NewRepository(db)
		getService    = getCategoriesController.NewService(getRepository)
		getHandler    = getCategoriesHandler.NewHandler(getService)
	)

	groupRoute := route.Group("/api/v1/categories")
	groupRoute.POST("/create", createHandler.CreateCategory)
	groupRoute.GET("/get", getHandler.GetAllCategories)
}
