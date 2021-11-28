package Store

import (
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/Utils"
	"CentralBankTask/internal/domain"
	"context"
	"time"
)

// BankStore structure of Bank database level
type BankStore struct {
	Conn Interface.ConnectionInterface
}

// AddDatesToBank adding dates to bankfilleddates
func (b2 BankStore) AddDatesToBank(ctx context.Context, currentDate, date time.Time) error {
	contextTransaction := context.Background()
	tx, err := b2.Conn.Begin(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.UpdateTransactionNotCreated,
		}
	}

	defer tx.Rollback(contextTransaction)

	_, err = tx.Exec(contextTransaction,
		"INSERT INTO bankfilleddates SELECT date_trunc('day', dd)::date FROM generate_series( $1::timestamp, $2::timestamp, '1 day'::interval) dd",
		currentDate, date)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.DateNotInserter,
		}
	}

	err = tx.Commit(contextTransaction)
	if err != nil {
		return &errPkg.Errors{
			Alias: errPkg.DateNotCommit,
		}
	}

	return nil
}

// CheckCurrentDate return the last date in bandfilleddates table
func (b2 BankStore) CheckCurrentDate(ctx context.Context, date time.Time) (time.Time, error) {
	var lastDate time.Time
	row := b2.Conn.QueryRow(ctx,
		"SELECT MAX(fulldate) FROM bankfilleddates")

	err := row.Scan(&lastDate)
	if err != nil {
		return lastDate, err
	}

	return lastDate, nil
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
		if err != nil {
			return false, downloadInterval, err
		}
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
		_, err := tx.Exec(contextTransaction,
			"UPDATE bankfilleddates SET isfilled = $1 WHERE fulldate = $2",
			true, convertedDate)
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
func (b2 BankStore) GetMaxValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error) {
	var maxValuteSlice []domain.ValuteValue
	row, err := b2.Conn.Query(ctx, "SELECT DISTINCT fulldate, valutename, valutevalue FROM bankinfo WHERE valutevalue = ( SELECT MAX(valutevalue) FROM bankinfo WHERE fulldate <= $1 AND fulldate > $2 ) AND fulldate <= $1 AND fulldate > $2 ORDER BY fulldate",
		dateInterval.BegD, dateInterval.EndD)

	if err != nil {
		return maxValuteSlice, err
	}

	for row.Next() {
		var maxValute domain.ValuteValue
		err = row.Scan(&maxValute.Date, &maxValute.Valute.Name, &maxValute.Valute.Value)
		if err != nil {
			return maxValuteSlice, err
		}

		maxValuteSlice = append(maxValuteSlice, maxValute)
	}

	return maxValuteSlice, nil
}

// GetMinValue info to get min value from bank
func (b2 BankStore) GetMinValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error) {
	var minValuteSlice []domain.ValuteValue
	row, err := b2.Conn.Query(ctx,
		"SELECT DISTINCT * FROM bankinfo WHERE valutevalue = ( SELECT MIN(valutevalue) FROM bankinfo WHERE fulldate <= $1 AND fulldate > $2 ) AND fulldate <= $1 AND fulldate > $2 ORDER BY fulldate",
		dateInterval.BegD, dateInterval.EndD)

	if err != nil {
		return minValuteSlice, err
	}

	for row.Next() {
		var minValute domain.ValuteValue
		err = row.Scan(&minValute.Date, &minValute.Valute.Name, &minValute.Valute.Value)
		if err != nil {
			return minValuteSlice, err
		}

		minValuteSlice = append(minValuteSlice, minValute)
	}

	return minValuteSlice, nil
}

// GetAverageValue info to get average value from bank
func (b2 BankStore) GetAverageValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error) {
	var avgValuteSlice []domain.ValuteValue
	row, err := b2.Conn.Query(ctx,
		"SELECT DISTINCT * FROM bankinfo WHERE valutevalue = (SELECT PERCENTILE_DISC(0.5) WITHIN GROUP(ORDER BY valutevalue) FROM bankinfo) AND fulldate <= $1 AND fulldate > $2",
		dateInterval.BegD, dateInterval.EndD)

	if err != nil {
		return avgValuteSlice, err
	}

	for row.Next() {
		var avgValute domain.ValuteValue
		err = row.Scan(&avgValute.Date, &avgValute.Valute.Name, &avgValute.Valute.Value)
		if err != nil {
			return avgValuteSlice, err
		}

		avgValuteSlice = append(avgValuteSlice, avgValute)
	}

	return avgValuteSlice, nil
}
