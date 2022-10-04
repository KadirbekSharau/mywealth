package registerAuthController

type InputUserRegister struct {
	Name string `json:"name"`
	Surname string `json:"surname"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}