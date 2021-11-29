package Interface

import (
	"CentralBankTask/inter/Bank"
	"context"
)

// BankInfoApplication implementation of user Application interface
type BankInfoApplication interface {
	SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error
	GetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) (Bank.ResponseBankInfoRequest, error)
}
