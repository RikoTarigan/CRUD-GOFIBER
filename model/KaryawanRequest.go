package model

type KaryawanCreateRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	Role     string `json:"role"`
	Password string `json:"password" validate:"required,min=6"`
}

type KaryawanUpdateRequest struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type KaryawanUpdateEmailRequest struct {
	Email string `json:"email" validate:"required"`
}