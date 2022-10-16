package handlers

import (
	"net/http"

	"github.com/KadirbekSharau/mywealth-backend/services/category"
	"github.com/KadirbekSharau/mywealth-backend/util"
	"github.com/gin-gonic/gin"
)

type CategoryHandler interface {
	CreateCategory(ctx *gin.Context)
	GetAllCategories(ctx *gin.Context)
}

type categoryHandler struct {
	service categoryService.Service
}

func NewCategoryHandler(service categoryService.Service) CategoryHandler {
	return &categoryHandler{service: service}
}

/* Get All Categories Handler */
func (h *categoryHandler) GetAllCategories(ctx *gin.Context) {

	fields, err := h.service.GetAllCategories()

	switch err {

	case "RESULTS_NOT_FOUND_404":
		util.APIResponse(ctx, "Data does not exist", http.StatusConflict, http.MethodPost, nil)

	default:
		util.APIResponse(ctx, "data found successfully", http.StatusOK, http.MethodPost, fields)
	}
}

func (h *categoryHandler) CreateCategory(ctx *gin.Context) {
	var input categoryService.InputCreateCategory
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

	_, errAccount := h.service.CreateCategory(&input)

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
