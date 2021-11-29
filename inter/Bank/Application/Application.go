package Application

import (
	"CentralBankTask/inter/Bank"
	"CentralBankTask/inter/Interface"
	"CentralBankTask/inter/Middleware/Error"
	"CentralBankTask/inter/Utils"
	"CentralBankTask/inter/domain"
	"context"
)

// BankApplication structure of bank application level
type BankApplication struct {
	BankStore Interface.BankInfoStore
}

// SetBankInfo setting bank information
func (b2 BankApplication) SetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error {
	timeOut := b.Date.Sub(Utils.GetTimeNowWithoutTime())
	if timeOut > 0 {
		return &Error.Errors{
			Alias: Error.NotValidDate,
			Text:  "There's no info about this day",
		}
	}

	lastDate, err := b2.BankStore.CheckCurrentDate(ctx, b.Date)
	if err != nil {
		return err
	}
	if b.Date.Sub(lastDate) > 0 {
		lastDate = lastDate.AddDate(0, 0, 1)
		err := b2.BankStore.AddDatesToBank(ctx, lastDate, b.Date)
		if err != nil {
			return err
		}
	}

	var dateStruct domain.DateInterval
	dateStruct = Utils.GetDate(b.Date)

	res, dateInterval, err := b2.BankStore.CheckDate(ctx, dateStruct)
	if err != nil {
		return err
	}

	for !res {
		var ValCurs []domain.ValCurs

		for _, date := range dateInterval.DateSlice {
			var download domain.ValCurs
			err = download.ReformatFile(Utils.CentralBankDataBaseURL+Utils.ConvertTimeToString(date, "/"), "smh.yml")
			if err != nil {
				return err
			}

			if Utils.ConvertStringToTime(download.Date) != date {
				download.Date = Utils.ConvertTimeToString(date, ".")
			}

			ValCurs = append(ValCurs, download)
		}
		err = b2.BankStore.UpdateBankInfo(ctx, ValCurs)
		if err != nil {
			return err
		}

		res, dateInterval, err = b2.BankStore.CheckDate(ctx, dateStruct)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetBankInfo returning business logic info
func (b2 BankApplication) GetBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) (Bank.ResponseBankInfoRequest, error) {
	var valuteSlice Bank.ResponseBankInfoRequest
	var dateStruct domain.DateInterval
	dateStruct = Utils.GetDate(b.Date)

	maxValue, err := b2.BankStore.GetMaxValue(ctx, dateStruct)
	if err != nil {
		return valuteSlice, err
	}

	valuteSlice.MaxValue = maxValue

	minValue, err := b2.BankStore.GetMinValue(ctx, dateStruct)
	if err != nil {
		return valuteSlice, err
	}

	valuteSlice.MinValue = minValue

	averageValue, err := b2.BankStore.GetAverageValue(ctx, dateStruct)
	if err != nil {
		return valuteSlice, err
	}

	valuteSlice.AverageValue = averageValue

	valuteSlice.Date = Utils.ConvertTimeToString(b.Date, ".")
	return valuteSlice, nil
}
