package models

type User struct {
	Id       int     `gorm:"primaryKey" json:"id"`
	Name     string  `gorm:"type:varchar(100)" json:"name"`
	Email    string  `gorm:"type:varchar(100);unique" json:"email"`
	Password string  `gorm:"type:text" json:"-"`
	Saldo    float64 `gorm:"type:decimal(10,2)" json:"saldo"`
}
