package getCategoriesHandler

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/controllers/categories/get"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetAllCategories(ctx *gin.Context)
}

type handler struct {
	service getCategoriesController.Service
}

func NewHandler(service getCategoriesController.Service) *handler {
	return &handler{service: service}
}

/* Get All Categories Handler */
func (h *handler) GetAllCategories(ctx *gin.Context) {

	fields, err := h.service.GetAllCategories()

	switch err {

	case "RESULTS_NOT_FOUND_404":
		util.APIResponse(ctx, "Data does not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "data found successfully", http.StatusOK, http.MethodPost, fields)
	}
}
