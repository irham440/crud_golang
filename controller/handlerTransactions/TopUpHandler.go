package handlertransactions

import (
	"belajar-go/dto"
	"belajar-go/utils"
	"encoding/json"
	"net/http"
)

func (s *ControllerTransaction) TopUpHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userId, ok1 := ctx.Value("user_id").(int)
	if !ok1 {
		utils.Error(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	var InputTopUp dto.TopUpSaldoRequest
	err := json.NewDecoder(r.Body).Decode(&InputTopUp)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	if InputTopUp.Amount <= 0 {
		utils.Error(w, http.StatusBadRequest, "amount must be greater than zero")
		return
	}

	url, err := s.service.CreateTransaction(ctx, InputTopUp.Amount, userId)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, "top up successful", map[string]string{"url": url})
}
