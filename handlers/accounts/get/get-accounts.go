package getAccountsHandler

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/KadirbekSharau/mywealth-backend/controllers/accounts/get"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllAccounts(ctx *gin.Context)
}

type handler struct {
	service getAccountsController.Service
}

func NewHandler(service getAccountsController.Service) *handler {
	return &handler{service: service}
}

/* Get All Accounts Handler */
func (h *handler) GetAllAcounts(ctx *gin.Context) {

	fields, err := h.service.GetAllAccounts()

	switch err {

	case "RESULTS_NOT_FOUND_404":
		util.APIResponse(ctx, "data does not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "data found successfully", http.StatusOK, http.MethodPost, fields)
	}
}