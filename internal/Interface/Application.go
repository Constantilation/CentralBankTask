package Interface

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/domain"
	"context"
)

// BankInfoApplication implementation of user Application interface
type BankInfoApplication interface {
	SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) (domain.ValCurs, error)
	MaxValue(ctx context.Context) (domain.ValCurs, error)
	MinValue(ctx context.Context) (domain.ValCurs, error)
	AverageValue(ctx context.Context) (domain.ValCurs, error)
}
