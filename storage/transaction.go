package storage

import (
	"belajar-go/models"
	"context"

	"gorm.io/gorm"
)

type Transaction interface {
	GetById(ctx context.Context, id int) ([]models.Transaction, error)
	CreateTransaction(ctx context.Context, db *gorm.DB, Transaction *models.Transaction) error
}

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (t *TransactionRepo) GetById(ctx context.Context, id int) ([]models.Transaction, error) {
	var transaction []models.Transaction
	err := t.db.WithContext(ctx).Where("user_id = ?", id).Find(&transaction).Error
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (t *TransactionRepo) CreateTransaction(ctx context.Context, db *gorm.DB, transaction *models.Transaction) error {
	err := db.WithContext(ctx).Create(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}
