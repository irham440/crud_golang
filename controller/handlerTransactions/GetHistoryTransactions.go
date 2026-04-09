package handlertransactions

import (
	"belajar-go/utils"
	"net/http"
)

func (s *ControllerTransaction) GetHistoryTransactionHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok := ctx.Value("user_id").(int)
	if !ok {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	transaction, err := s.service.GetHistoryTransaction(ctx, userId)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "transaction history retrieved successfully", transaction)
}
