package API

import (
	"CentralBankTask/internal/Bank"
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
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
	date := c.Param("date")
	ctx := c.Request().Context()
	updateBankStruct := Bank.UpdateBankInfoRequest{
		Date: date,
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

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, handler Interface.BankInfoAPI) {
	e.GET("/:date", handler.GetBankInfoHandler)
}
