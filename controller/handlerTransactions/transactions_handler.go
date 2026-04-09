package handlertransactions

import (
	"belajar-go/services/transactions"
)

type ControllerTransaction struct {
	service transactions.TransactionService
}

func NewControllerTransaction(service transactions.TransactionService) *ControllerTransaction {
	return &ControllerTransaction{service: service}
}
