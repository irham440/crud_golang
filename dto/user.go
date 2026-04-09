package dto

type getProfileRequest struct {
	Id int `json:"id"`
}

type getprofileResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}
type CreateUserRequest struct {
	Name     string `json:"name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=abcdefghijklmnopqrstuvwxyz"`
}

type createUserResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}

type TopUpSaldoRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	Id     int     `json:"id" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,alphanum,contains=1"`
}
