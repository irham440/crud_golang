package storage

import (
	"belajar-go/models"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(id int) (models.User, error)
	Save(user models.User) error
	GetEmail(email string) (string, error)
}

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) GetById(id int) (models.User, error) {
	var user models.User
	err := p.db.First(&user, id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (p *PostgresRepo) Save(user models.User) error {
	return p.db.Create(&user).Error
}

func (p *PostgresRepo) GetEmail(email string) (string, error) {
	var foundEmail string
	err := p.db.Model(models.User{}).Select("email").Where("email = ?", email).First(&foundEmail).Error
	if err != nil {
		return "", err
	}
	fmt.Println(foundEmail)
	return foundEmail, nil
}
