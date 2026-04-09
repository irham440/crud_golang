package transactions

import (
	"belajar-go/models"
	"belajar-go/services/email"
	"belajar-go/storage"

	"context"
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type TransactionService interface {
	GetHistoryTransaction(ctx context.Context, id int) ([]models.Transaction, error)
	TopUpSaldo(ctx context.Context, id int, amount float64) error

	HandleMidtransNotification(ctx context.Context, id int, amount float64) error
	CreateTransaction(ctx context.Context, amount float64, id int) (string, error)
}

type ITransactionService struct {
	tx     storage.Transaction
	user   storage.UserRepository
	db     *gorm.DB
	client *http.Client
	Email  email.EmailService
}

func NewITransactionService(repo storage.Transaction, user storage.UserRepository, db *gorm.DB, Email email.EmailService) *ITransactionService {
	return &ITransactionService{tx: repo, user: user, db: db, client: &http.Client{}, Email: Email}
}

func (t *ITransactionService) GetHistoryTransaction(ctx context.Context, id int) ([]models.Transaction, error) {
	transaction, err := t.tx.GetById(ctx, id)
	if err != nil {
		return nil, errors.New("Gagal ambil data transaksi")
	}
	return transaction, nil
}
