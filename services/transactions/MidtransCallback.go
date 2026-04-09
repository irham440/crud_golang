package transactions

import (
	"belajar-go/models"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (t *ITransactionService) HandleMidtransNotification(ctx context.Context, id int, amount float64) error {
	_, err := t.user.GetById(ctx, id)
	if err != nil {
		return errors.New("user not found")
	}
	err = t.db.Transaction(func(tx *gorm.DB) error {
		err = t.user.AddSaldo(ctx, tx, id, amount)
		if err != nil {
			return err
		}
		Transaction := &models.Transaction{
			UserId: id,
			Amount: amount,
			Status: "Topup",
		}
		fmt.Println("Midtrans notification received for user ID:", id, "with amount:", amount)
		err = t.tx.CreateTransaction(ctx, tx, Transaction)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return errors.New("failed to top up saldo")
	}
	t.Email.SendEmail("user@example.com", "Top Up Saldo", "Anda telah melakukan top up saldo sebesar $amount")
	return nil
}
