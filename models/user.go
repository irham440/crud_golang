package models

type User struct {
	Id       int     `gorm:"primaryKey" json:"id"`
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"-"`
	Saldo    float64 `json:"saldo"`
}
