package accountService

type InputAccountsByUserID struct {
	UserID uint   `json:"user_id" validate:"required"`
}

// InputUpdateAccountByID is an input for updating account by user id
type InputUpdateAccountByID struct {
	Name   string `json:"name"`
	ID uint   `json:"id"`
}

// InputCreateAccount is an input for creating an account
type InputCreateAccount struct {
	Name   string `json:"name"`
	UserID uint   `json:"user_id" validate:"required"`
}