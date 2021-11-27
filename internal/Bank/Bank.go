package Bank

import "CentralBankTask/internal/domain"

// UpdateBankInfoRequest represent struct for update bank info for last 90 days
type UpdateBankInfoRequest struct {
	Date string `json:"date"`
}

type ResponseBankInfoRequest struct {
	Date   string          `json:"date"`
	Valute []domain.Valute `json:"valute"`
}
