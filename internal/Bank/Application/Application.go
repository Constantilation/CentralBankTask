package Application

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	"CentralBankTask/internal/Utils"
	"CentralBankTask/internal/domain"
	"context"
)

// BankApplication structure of bank application level
type BankApplication struct {
	BankStore Interface.BankInfoStore
}

// SetBankInfo setting bank information
func (b2 BankApplication) SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error {
	var dateStruct domain.DateInterval
	dateStruct = Utils.GetDate(b.Date)

	res, dateInterval, err := b2.BankStore.CheckDate(ctx, dateStruct)
	if err != nil {
		return err
	}

	if !res {
		var ValCurs []domain.ValCurs

		for _, value := range dateInterval.DateSlice {
			var download domain.ValCurs
			err = download.ReformatFile(Utils.CentralBankDataBaseURL+Utils.ConvertTimeToString(value), "smh.yml")
			if err != nil {
				return err
			}
			ValCurs = append(ValCurs, download)
		}
		err = b2.BankStore.UpdateBankInfo(ctx, ValCurs)
		if err != nil {
			return err
		}

		return nil
	}

	return nil
}

// GetBankInfo returning business logic info
func (b2 BankApplication) GetBankInfo(ctx context.Context) (Bank.ResponseBankInfoRequest, error) {
	//var bankInfo Bank.ResponseBankInfoRequest
	panic(1)
}
