package Application

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	"CentralBankTask/internal/domain"
	"context"
)

// BankApplication structure of bank application level
type BankApplication struct {
	BankStore Interface.BankInfoStore
}

// SetBankInfo setting bank information
func (b2 BankApplication) SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// MaxValue getting max value
func (b2 BankApplication) MaxValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// MinValue getting min value
func (b2 BankApplication) MinValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}

// AverageValue getting average value
func (b2 BankApplication) AverageValue(ctx context.Context) (domain.ValCurs, error) {
	//TODO implement me
	panic("implement me")
}
