package routes

import (
	loginAuthController "github.com/KadirbekSharau/mywealth-backend/controllers/auth/login"
	registerAuthController "github.com/KadirbekSharau/mywealth-backend/controllers/auth/register"
	LoginHandler "github.com/KadirbekSharau/mywealth-backend/handlers/auth/login"
	registerHandler "github.com/KadirbekSharau/mywealth-backend/handlers/auth/register"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* @description All Auth routes */
func InitAuthRoutes(db *gorm.DB, route *gin.Engine) {
	var (
		loginRepository loginAuthController.Repository = loginAuthController.NewRepositoryLogin(db)
		loginService loginAuthController.Service = loginAuthController.NewServiceLogin(loginRepository)
		loginHandler LoginHandler.Handler = LoginHandler.NewLoginHandler(loginService)

		registerRepository registerAuthController.Repository = registerAuthController.NewRepositoryRegister(db)
		registerService registerAuthController.Service = registerAuthController.NewServiceRegister(registerRepository)
		registerHandler registerHandler.Handler = registerHandler.NewHandlerRegister(registerService)
	)

	groupRoute := route.Group("/api/v1/auth")
	groupRoute.POST("/user/login", loginHandler.UserLoginHandler)
	// groupRoute.POST("/admin/login", loginHandler.AdminLoginHandler)
	groupRoute.POST("/user/register", registerHandler.ActiveUserRegisterHandler)
}