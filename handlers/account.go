package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/services/account"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	UpdateAccountByID(ctx *gin.Context)
	GetAllAcounts(ctx *gin.Context)
	CreateAccount(ctx *gin.Context) 
}

type accountHandler struct {
	service accountService.Service
}

func NewAccountHandler(service accountService.Service) *accountHandler {
	return &accountHandler{service: service}
}

/* Get All Accounts Handler */
func (h *accountHandler) GetAllAcounts(ctx *gin.Context) {

	fields, err := h.service.GetAllAccounts()

	switch err {

	case "RESULTS_NOT_FOUND_404":
		util.APIResponse(ctx, "data does not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "data found successfully", http.StatusOK, http.MethodPost, fields)
	}
}

// CreateAccount creates the account
func (h *accountHandler) CreateAccount(ctx *gin.Context) {
	var input accountService.InputCreateAccount
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

	_, errAccount := h.service.CreateAccount(&input)

	switch errAccount {
	case "CREATE_CONFLICT_409":
		util.APIResponse(ctx, "Name field already exist", http.StatusConflict, http.MethodPost, nil)
		return

	case "CREATE_FAILED_403":
		util.APIResponse(ctx, "Create new instance failed", http.StatusForbidden, http.MethodPost, nil)
		return

	default:
		util.APIResponse(ctx, "Create new instance successfully", http.StatusCreated, http.MethodPost, nil)
	}
}

func (h *accountHandler) UpdateAccountByID(ctx *gin.Context) {}