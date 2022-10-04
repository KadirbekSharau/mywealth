package registerHandler

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/controllers/auth/register"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	ActiveUserRegisterHandler(ctx *gin.Context)
}

type handler struct {
	service registerAuthController.Service
}

func NewHandlerRegister(service registerAuthController.Service) *handler {
	return &handler{service: service}
}

/* Active User Register Handler */
func (h *handler) ActiveUserRegisterHandler(ctx *gin.Context) {

	var input registerAuthController.InputUserRegister
	ctx.ShouldBindJSON(&input)
	errResponse, errCount := util.GoValidator(input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ActiveUserRegisterService(&input)
	ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}