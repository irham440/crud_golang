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
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type createUserResponse struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Email string  `json:"email"`
	Saldo float64 `json:"saldo"`
}

type TopUpSaldoRequest struct {
	Id     int     `json:"id"`
	Amount float64 `json:"amount"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
