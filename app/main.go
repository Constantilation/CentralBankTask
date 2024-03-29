package main

import (
	build "CentralBankTask/build"
	"CentralBankTask/config"
	templates "CentralBankTask/files/template"
	"CentralBankTask/inter/Bank/API"
	"CentralBankTask/inter/Interface"
	"CentralBankTask/inter/Middleware"
	errors "CentralBankTask/inter/Middleware/Error"
	utils "CentralBankTask/inter/Utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.uber.org/zap"
	"html/template"
	"os"
)

func runServer() {
	var logger utils.Logger
	logger.Log = utils.NewLogger("logs.txt")

	defer func(loggerErrWarn errors.MultiLogger) {
		errLogger := loggerErrWarn.Sync()
		if errLogger != nil {
			zap.S().Errorf("LoggerErrWarn the buffer could not be cleared %v", errLogger)
			os.Exit(4)
		}
	}(logger.Log)

	errConfig, configStructure := build.InitConfig()
	if errConfig != nil {
		logger.Log.Errorf("%s", errConfig.Error())
		return
	}
	dbConfig := configStructure[0].(config.DBConfig)
	appConfig := configStructure[1].(config.AppConfig)

	connectionJSON, err := build.CreateDb(dbConfig.Db, appConfig.Primary.Debug)
	if err != nil {
		logger.Log.Errorf("Unable to connect to database: %s", err.Error())
		os.Exit(2)
	}

	startStructure := build.SetUp(connectionJSON, logger.Log)
	BankInfo := startStructure[0].(Interface.BankInfoAPI)

	t := &templates.Template{
		Templates: template.Must(template.ParseGlob("./index.html")),
	}
	e := echo.New()
	e.Renderer = t
	e.Pre(middleware.AddTrailingSlash())
	middl := Middleware.InitMiddleware()
	logInfo := Middleware.InfoMiddleware{
		Logger: logger.Log,
	}
	e.Use(middl.CORS)
	e.Use(logInfo.LogURL)
	API.NewBankHandler(e, BankInfo)
	err = e.Start(appConfig.Port)
	if err != nil {
		logger.Log.Errorf("Listen and server error: %v", err)
		os.Exit(3)
	}
}

func main() {
	runServer()
}
