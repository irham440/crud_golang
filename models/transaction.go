package models

type Transaction struct {
	Id     int     `gorm:"primaryKey" json:"id"`
	UserId int     `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string  `json:"status"`
}
