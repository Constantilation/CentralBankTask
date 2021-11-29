package Bank

import (
	"CentralBankTask/inter/domain"
	"time"
)

// UpdateBankInfoRequest represent struct for update bank info for last 90 days
type UpdateBankInfoRequest struct {
	Date time.Time `json:"date"`
}

type ResponseBankInfoRequest struct {
	Date         string               `json:"Request_Data"`
	MaxValue     []domain.ValuteValue `json:"Max_Valute"`
	MinValue     []domain.ValuteValue `json:"Min_Valute"`
	AverageValue []domain.ValuteValue `json:"Average_Valute"`
}
