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

// CheckDate checking if info of this date is already exist. Return false, if date is not exist
func (b2 BankStore) CheckDate(ctx context.Context, date string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

// UpdateBankInfo method to update bank info
func (b2 BankStore) UpdateBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error {
	//TODO implement me
	panic("implement me")
}

// GetMaxValue info to get max value from bank
func (b2 BankStore) GetMaxValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// GetMinValue info to get min value from bank
func (b2 BankStore) GetMinValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// GetAverageValue info to get average value from bank
func (b2 BankStore) GetAverageValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}
