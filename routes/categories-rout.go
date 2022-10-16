package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/services/category"
	"github.com/KadirbekSharau/mywealth-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Categories routes */
func InitCategoryRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = categoryService.NewRepository(db)
		service    = categoryService.NewService(repository)
		handler    = handlers.NewCategoryHandler(service)
	)

	groupRoute := route.Group("/api/v1/categories")
	groupRoute.POST("/create", handler.CreateCategory)
	groupRoute.GET("/get", handler.GetAllCategories)
}
