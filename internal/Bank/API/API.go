package API

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/Utils"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

// BankAPI structure of Bank API
type BankAPI struct {
	BankApplication Interface.BankInfoApplication
	Logger          errPkg.MultiLogger
}

// GetBankInfoHandler implementation of getting info
func (b BankAPI) GetBankInfoHandler(c echo.Context) error {
	day, _ := Utils.InterfaceConvertInt(c.Param("day"))
	month, _ := Utils.InterfaceConvertInt(c.Param("month"))
	year, _ := Utils.InterfaceConvertInt(c.Param("year"))
	data := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	ctx := c.Request().Context()
	updateBankStruct := Bank.UpdateBankInfoRequest{
		Date: data,
	}

	err := b.BankApplication.SetBankInfo(ctx, &updateBankStruct)

	checkError := &errPkg.CheckError{
		Logger: b.Logger,
	}

	if err != nil {
		return c.JSON(checkError.CheckErrorsBank(err),
			errPkg.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: err.Error(),
			})
	}

	bankInfo, err := b.BankApplication.GetBankInfo(ctx)
	if err != nil {
		return c.JSON(checkError.CheckErrorsBank(err),
			errPkg.ResultError{
				Status:  http.StatusInternalServerError,
				Explain: err.Error(),
			})
	}

	return c.JSON(http.StatusOK, bankInfo)
}

// NewBankHandler will initialize the articles/ resources endpoint
func NewBankHandler(e *echo.Echo, handler Interface.BankInfoAPI) {
	ug := e.Group("/:day/:month/:year")
	ug.GET("/", handler.GetBankInfoHandler)
}
