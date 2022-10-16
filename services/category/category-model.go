package categoryService

type InputCreateCategory struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}