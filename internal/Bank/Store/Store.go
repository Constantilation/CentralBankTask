package Store

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	"CentralBankTask/internal/domain"
	"context"
)

// BankStore structure of Bank database level
type BankStore struct {
	Conn Interface.ConnectionInterface
}

// UpdateBankInfo method to update bank info
func (b2 BankStore) UpdateBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error {
	//TODO implement me
	panic("implement me")
}

// MaxValue info to get max value from bank
func (b2 BankStore) MaxValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// MinValue info to get min value from bank
func (b2 BankStore) MinValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// AverageValue info to get average value from bank
func (b2 BankStore) AverageValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}
