package Interface

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/domain"
	"context"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx"
)

// ConnectionInterface implementation of database methods interface
type ConnectionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

// BankInfoStore implementation of Store interface methods
type BankInfoStore interface {
	UpdateBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error
	CheckDate(ctx context.Context, date string) (bool, error)
	GetMaxValue(ctx context.Context) (domain.ValCurs, error)
	GetMinValue(ctx context.Context) (domain.ValCurs, error)
	GetAverageValue(ctx context.Context) (domain.ValCurs, error)
}
