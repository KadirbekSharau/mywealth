package routes

import (
	"github.com/KadirbekSharau/mywealth-backend/services/auth"
	"github.com/KadirbekSharau/mywealth-backend/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		repository = authService.NewRepository(db)
		service    = authService.NewService(repository)
		authHandler    = handlers.NewAuthHandler(service)
	)

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/user/login", authHandler.UserLogin)
	// groupRoute.POST("/admin/login", loginHandler.AdminLoginHandler)
	groupRoute.POST("/user/register", authHandler.ActiveUserRegister)
}
