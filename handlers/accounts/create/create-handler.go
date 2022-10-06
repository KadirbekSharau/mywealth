package createAccountHandler

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/controllers/accounts/create"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateAccountHandler(ctx *gin.Context) 
}

type handler struct {
	service createAccountController.Service
}

func NewCreateAccountHandler(service createAccountController.Service) *handler {
	return &handler{service: service}
}

func (h *handler) CreateAccountHandler(ctx *gin.Context) {
	var input createAccountController.InputCreateAccount
	ctx.ShouldBindJSON(&input)

	config := util.ErrorConfig{
		Options: []util.ErrorMetaConfig{
			{
				Tag:     "required",
				Field:   "Name",
				Message: "name is required on body",
			},
			{
				Tag:     "required",
				Field:   "UserID",
				Message: "user id is required on body",
			},
		},
	}

	errResponse, errCount := util.GoValidator(&input, config.Options)

	if errCount > 0 {
		util.ValidatorErrorResponse(ctx, http.StatusBadRequest, http.MethodPost, errResponse)
		return
	}

	_, errAccount := h.service.CreateAccountService(&input)

	switch errAccount {
	case "CREATE_ACCOUNT_CONFLICT_409":
		util.APIResponse(ctx, "Name field already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_ACCOUNT_FAILED_403":
		util.APIResponse(ctx, "Create new account failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new account successfully", http.StatusCreated, http.MethodPost, nil)
	}
}