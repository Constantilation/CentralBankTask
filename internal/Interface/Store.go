package Interface

import (
	"CentralBankTask/internal/domain"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"time"
)

// ConnectionInterface implementation of database methods interface
type ConnectionInterface interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

// BankInfoStore implementation of Store interface methods
type BankInfoStore interface {
	CheckCurrentDate(ctx context.Context, date time.Time) (time.Time, error)
	AddDatesToBank(ctx context.Context, currentDate, date time.Time) error
	UpdateBankInfo(ctx context.Context, valuteData []domain.ValCurs) error
	CheckDate(ctx context.Context, date domain.DateInterval) (bool, domain.DownloadInterval, error)
	GetMaxValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error)
	GetMinValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error)
	GetAverageValue(ctx context.Context, dateInterval domain.DateInterval) ([]domain.ValuteValue, error)
}
