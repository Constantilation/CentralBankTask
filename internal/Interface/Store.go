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

// TransactionInterface implementation of database transaction methods interface
type TransactionInterface interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	LargeObjects() pgx.LargeObjects
	Prepare(ctx context.Context, name, sql string) (*pgconn.StatementDescription, error)
	QueryFunc(ctx context.Context, sql string, args []interface{}, scans []interface{}, f func(pgx.QueryFuncRow) error) (pgconn.CommandTag, error)
	Conn() *pgx.Conn
}

// BankInfoStore implementation of Store interface methods
type BankInfoStore interface {
	UpdateBankInfo(ctx context.Context, b *Bank.UpdateBankInfoRequest) error
	MaxValue(ctx context.Context) (domain.ValCurs, error)
	MinValue(ctx context.Context) (domain.ValCurs, error)
	AverageValue(ctx context.Context) (domain.ValCurs, error)
}
