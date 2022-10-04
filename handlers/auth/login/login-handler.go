package loginHandler

import (
	"net/http"

	loginAuth "github.com/KadirbekSharau/mywealth-backend/controllers/auth/login"
	util "github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	UserLoginHandler(ctx *gin.Context)
}

type handler struct {
	service loginAuth.Service
}

func NewLoginHandler(service loginAuth.Service) *handler {
	return &handler{service: service}
}

/* User Login Handler */
func (h *handler) UserLoginHandler(ctx *gin.Context) {

	var input loginAuth.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.UserLoginService(&input)

	UserLoginTokenHandler(ctx, errLogin, resultLogin)
}