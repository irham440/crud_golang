package handlertransactions

import (
	"belajar-go/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (s *ControllerTransaction) HandleMidtransNotification(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var notification struct {
		TransactionStatus string `json:"transaction_status"`
		GrossAmount       string `json:"gross_amount"`
		Metadata          struct {
			IDUser   int     `json:"idUser"`
			Price    float64 `json:"price"`
			Quantity int     `json:"quantity"`
			Name     string  `json:"name"`
		} `json:"metadata"`
	}

	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		utils.Error(w, http.StatusBadRequest, "invalid request body")
		return
	}

	userID := notification.Metadata.IDUser

	totalAmount, _ := strconv.ParseFloat(notification.GrossAmount, 64)

	fmt.Printf("Processing: UserID %d, Amount %.2f, Status %s\n",
		userID, totalAmount, notification.TransactionStatus)

	if notification.TransactionStatus == "settlement" || notification.TransactionStatus == "capture" {
		err = s.service.HandleMidtransNotification(ctx, userID, totalAmount)
		if err != nil {
			fmt.Printf("Service Error: %v\n", err)
			utils.Error(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	utils.Success(w, http.StatusOK, "notification processed successfully", nil)
}
