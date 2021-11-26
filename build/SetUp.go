package build

import (
	"CentralBankTask/config"
	"CentralBankTask/internal/Bank/API"
	"CentralBankTask/internal/Bank/Application"
	"CentralBankTask/internal/Bank/Store"
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"context"
	"github.com/jackc/pgx/pgxpool"
	"github.com/spf13/viper"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfNameURLS = "urls"
	ConfType     = "yml"
	ConfPath     = "./config"
)

// InitConfig function to initialize structures
func InitConfig() (error, []interface{}) {
	viper.AddConfigPath(ConfPath)
	viper.SetConfigType(ConfType)

	viper.SetConfigName(ConfNameMain)
	errRead := viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	appConfig := config.AppConfig{}
	errUnmarshal := viper.Unmarshal(&appConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	viper.SetConfigName(ConfNameDB)
	errRead = viper.ReadInConfig()
	if errRead != nil {
		return &errPkg.Errors{
			Alias: errRead.Error(),
		}, nil
	}
	dbConfig := config.DBConfig{}
	errUnmarshal = viper.Unmarshal(&dbConfig)
	if errUnmarshal != nil {
		return &errPkg.Errors{
			Alias: errUnmarshal.Error(),
		}, nil
	}

	var result []interface{}
	result = append(result, dbConfig)
	result = append(result, appConfig)

	return nil, result
}

// SetUp setting up user essence
func SetUp(connectionDB Interface.ConnectionInterface, logger errPkg.MultiLogger) []interface{} {
	BankStore := Store.BankStore{Conn: connectionDB}
	BankApp := Application.BankApplication{BankStore: &BankStore}
	BankInfo := API.BankAPI{
		BankApplication: &BankApp,
		Logger:          logger,
	}

	var _ Interface.BankInfoAPI = &BankInfo

	var result []interface{}
	result = append(result, BankInfo)

	return result
}

// CreateDb Creating data base structure
func CreateDb(configDB config.Database) (*pgxpool.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+configDB.UserName+":"+configDB.Password+
			"@"+configDB.Host+":"+configDB.Port+"/"+configDB.SchemaName)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
		}
	}
	return conn, nil
}
