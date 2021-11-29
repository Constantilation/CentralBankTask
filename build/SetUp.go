package build

import (
	"CentralBankTask/config"
	"CentralBankTask/inter/Bank/API"
	"CentralBankTask/inter/Bank/Application"
	"CentralBankTask/inter/Bank/Store"
	"CentralBankTask/inter/Interface"
	errPkg "CentralBankTask/inter/Middleware/Error"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	pgxpool2 "github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

const (
	ConfNameMain = "main"
	ConfNameDB   = "database"
	ConfType     = "yml"
	ConfPath     = "."
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
func CreateDb(configDB config.Database, debug bool) (*pgxpool2.Pool, error) {
	var err error
	conn, err := pgxpool.Connect(context.Background(),
		"postgres://"+configDB.UserName+":"+configDB.Password+
			"@"+configDB.Host+":"+configDB.Port+"/"+configDB.SchemaName)
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBNotConnect,
			Text:  err.Error(),
		}
	}

	if debug {
		file, err := ioutil.ReadFile("./DeleteTables.sql")
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBDeleteFileNotFound,
			}
		}

		requests := strings.Split(string(file), ";")
		for _, request := range requests {
			_, err = conn.Exec(context.Background(), request)
			if err != nil {
				return nil, &errPkg.Errors{
					Alias: errPkg.MCreateDBNotDeleteTables,
				}
			}
		}
	}
	file, err := ioutil.ReadFile("./CreateTables.sql")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBCreateFileNotFound,
		}
	}

	requests := strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBNotCreateTables,
			}
		}
	}

	file, err = ioutil.ReadFile("./Fill.sql")
	if err != nil {
		return nil, &errPkg.Errors{
			Alias: errPkg.MCreateDBFillFileNotFound,
		}
	}

	requests = strings.Split(string(file), ";")
	for _, request := range requests {
		_, err = conn.Exec(context.Background(), request)
		if err != nil {
			return nil, &errPkg.Errors{
				Alias: errPkg.MCreateDBNotFillTables,
			}
		}
	}

	return conn, nil
}
