package Application

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	"context"
)

// BankApplication structure of bank application level
type BankApplication struct {
	BankStore Interface.BankInfoStore
}

// SetBankInfo setting bank information
func (b2 BankApplication) SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error {
	res, err := b2.BankStore.CheckDate(ctx, b.Date)
	if err != nil {
		return err
	}

	if !res {
		err = b2.BankStore.UpdateBankInfo(ctx, b)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

// GetBankInfo returning business logic info
func (b2 BankApplication) GetBankInfo(ctx context.Context) (Bank.ResponseBankInfoRequest, error) {
	var bankInfo Bank.ResponseBankInfoRequest


}
