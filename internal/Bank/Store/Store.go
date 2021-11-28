package Store

import (
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/Utils"
	"CentralBankTask/internal/domain"
	"context"
	"fmt"
	"time"
)

// BankStore structure of Bank database level
type BankStore struct {
	Conn Interface.ConnectionInterface
}

// CheckDate checking if info of this date is already exist. Return false, if date is not exist
func (b2 BankStore) CheckDate(ctx context.Context, date domain.DateInterval) (bool, domain.DownloadInterval, error) {
	var downloadInterval domain.DownloadInterval
	row, err := b2.Conn.Query(ctx,
		"SELECT fulldate FROM bankfilleddates WHERE fulldate <= $1 AND fulldate > $2 AND isfilled IS FALSE", date.BegD, date.EndD)

	if err != nil {
		return false, downloadInterval, err
	}

	for row.Next() {
		var fulldate time.Time
		err = row.Scan(&fulldate)
		downloadInterval.DateSlice = append(downloadInterval.DateSlice, fulldate)
	}

	if len(downloadInterval.DateSlice) > 0 {
		return false, downloadInterval, nil
	}
	return true, downloadInterval, nil
}

// UpdateBankInfo method to update bank info
func (b2 BankStore) UpdateBankInfo(ctx context.Context, valuteData []domain.ValCurs) error {
	contextTransaction := context.Background()
	tx, err := b2.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.UpdateTransactionNotCreated,
		}
	}

	defer tx.Rollback(contextTransaction)

	for _, date := range valuteData {
		convertedDate := Utils.ConvertStringToTime(date.Date)
		for _, valute := range date.Valute {
			_, err = tx.Exec(contextTransaction,
				"INSERT INTO bankinfo (fulldate, valutename, valutevalue) VALUES ($1, $2, $3)",
				convertedDate, valute.Name, valute.Value)
			if err != nil {
				return &errPkg.Errors{
					Alias: errPkg.ValutesNotInsert,
				}
			}
		}
		result, err := tx.Exec(contextTransaction,
			"UPDATE bankfilleddates SET isfilled = $1 WHERE fulldate = $2",
			true, convertedDate)
		fmt.Println(result)
		if err != nil {
			return &errPkg.Errors{
				Alias: errPkg.ValutesNotInsert,
			}
		}
	}
	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.ValutesNotCommit,
		}
	}

	return nil
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
