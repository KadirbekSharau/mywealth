package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/handlers/helpers"
	"github.com/KadirbekSharau/mywealth-backend/services/auth"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	UserLogin(ctx *gin.Context)
	ActiveUserRegister(ctx *gin.Context)
}

type authHandler struct {
	service authService.Service
}

func NewAuthHandler(service authService.Service) AuthHandler {
	return &authHandler{service: service}
}

/* User Login Handler */
func (h *authHandler) UserLogin(ctx *gin.Context) {

	var input authService.InputLogin
	ctx.ShouldBindJSON(&input)

	errResponse, errCount := util.GoValidator(&input, helpers.Config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultLogin, errLogin := h.service.UserLogin(&input)

	helpers.UserLoginTokenHandler(ctx, errLogin, resultLogin)
}

/* Active User Register Handler */
func (h *authHandler) ActiveUserRegister(ctx *gin.Context) {

	var input authService.InputUserRegister
	ctx.ShouldBindJSON(&input)
	conf := helpers.Config.Options
	conf = append(conf, util.ErrorMetaConfig{
		Tag:     "gte",
		Field:   "Password",
		Message: "password minimum must be 8 character",
	},)
	errResponse, errCount := util.GoValidator(input, conf)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	resultRegister, errRegister := h.service.ActiveUserRegister(&input)
	helpers.ErrUserRegisterHandler(resultRegister, ctx, errRegister)
}
