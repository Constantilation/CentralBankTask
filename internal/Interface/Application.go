package Interface

import (
	"CentralBankTask/internal/Bank"
	"context"
)

// BankInfoApplication implementation of user Application interface
type BankInfoApplication interface {
	SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error
	GetBankInfo(ctx context.Context) (Bank.ResponseBankInfoRequest, error)
}
