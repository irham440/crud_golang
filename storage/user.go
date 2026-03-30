package storage

import (
	"belajar-go/models"
	"context"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetById(ctx context.Context, id int) (*models.User, error)
	Save(ctx context.Context, user *models.User) error
	GetEmail(ctx context.Context, email string) (string, error)
	Update(ctx context.Context, user *models.User) error
	AddSaldo(ctx context.Context, tx *gorm.DB, id int, amount float64) error
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type PostgresRepo struct {
	db *gorm.DB
}

func NewPostgresRepo(db *gorm.DB) *PostgresRepo {
	return &PostgresRepo{db: db}
}

func (p *PostgresRepo) GetById(ctx context.Context, id int) (*models.User, error) {
	var user models.User
	err := p.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PostgresRepo) Save(ctx context.Context, user *models.User) error {
	err := p.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) GetEmail(ctx context.Context, email string) (string, error) {
	var foundEmail string
	err := p.db.Model(models.User{}).Select("email").Where("email = ?", email).First(&foundEmail).Error
	if err != nil {
		return "", err
	}
	return foundEmail, nil
}

func (p *PostgresRepo) Update(ctx context.Context, user *models.User) error {
	err := p.db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) AddSaldo(ctx context.Context, tx *gorm.DB, id int, amount float64) error {
	err := tx.Model(&models.User{}).Where("id = ?", id).Update("saldo", gorm.Expr("saldo + ?", amount)).Error
	if err != nil {
		return err
	}
	return nil
}

func (p *PostgresRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := p.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
