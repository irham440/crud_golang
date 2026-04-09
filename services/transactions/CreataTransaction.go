package transactions

import (
	"belajar-go/dto"

	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

func (t *ITransactionService) CreateTransaction(ctx context.Context, amount float64, id int) (string, error) {

	reqBody := dto.TopUpSaldoRequest{
		Amount: amount,
		Id:     id,
	}
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", errors.New("Gagal membuat transaksi")
	}

	req, err := http.NewRequest("POST", "http://localhost:3000/top-up", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", errors.New("Gagal membuat transaksi request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "rahasia_aman_123")

	resp, err := t.client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		return "", errors.New("Gagal membuat transaksi")
	}
	defer resp.Body.Close()
	type Response struct {
		Message string `json:"message"`
		Success bool   `json:"success"`
		Url     string `json:"url"`
	}
	var response Response
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil || !response.Success {
		return "", errors.New("Gagal membuat transaksi")
	}
	return response.Url, nil
}
