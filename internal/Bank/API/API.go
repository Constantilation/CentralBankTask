package API

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"CentralBankTask/internal/Utils"
	"CentralBankTask/internal/domain"
	"github.com/labstack/echo"
	"net/http"
)

// BankAPI structure of Bank API
type BankAPI struct {
	BankApplication Interface.BankInfoApplication
	Logger          errPkg.MultiLogger
}

// GetBankInfoHandler implementation of getting info
func (b BankAPI) GetBankInfoHandler(c echo.Context) error {
	var data domain.Date
	data.DD, _ = Utils.InterfaceConvertInt(c.Param("day"))
	data.MM, _ = Utils.InterfaceConvertInt(c.Param("month"))
	data.YY, _ = Utils.InterfaceConvertInt(c.Param("year"))
	ctx := c.Request().Context()
	updateBankStruct := Bank.UpdateBankInfoRequest{
		Date: Utils.ConvertDateToString(data),
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
