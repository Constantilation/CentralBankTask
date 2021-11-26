package main

import (
	build "CentralBankTask/build"
	"CentralBankTask/config"
	"CentralBankTask/internal/Bank/API"
	"CentralBankTask/internal/Interface"
	"CentralBankTask/internal/Middleware"
	errors "CentralBankTask/internal/Middleware/Error"
	utils "CentralBankTask/internal/Utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"os"
)

func runServer() {
	var logger utils.Logger
	logger.Log = utils.NewLogger("logs.txt")

	defer func(loggerErrWarn errors.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(1)
		}
	}(logger.Log)

	errConfig, configStructure := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		return
	}
	dbConfig := configStructure[0].(config.DBConfig)
	appConfig := configStructure[1].(config.AppConfig)

	connectionJSON, err := build.CreateDb(dbConfig.Db)
	if err != nil {
		logger.Log.Errorf("Unable to connect to database: %s", err.Error())
		os.Exit(1)
	}

	startStructure := build.SetUp(connectionJSON, logger.Log)
	userInfo := startStructure[0].(Interface.BankInfoAPI)

	e := echo.New()
	e.Pre(middleware.AddTrailingSlash())
	middl := Middleware.InitMiddleware()
	logInfo := Middleware.InfoMiddleware{
		Logger: logger.Log,
	}
	e.Use(middl.CORS)
	e.Use(logInfo.LogURL)
	API.NewUserHandler(e, userInfo)

	err = e.Start(appConfig.Port)
	if err != nil {
		logger.Log.Errorf("Listen and server error: %v", err)
		os.Exit(1)
	}
}

func main() {
	runServer()
}
