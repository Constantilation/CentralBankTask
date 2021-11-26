package API

import (
	"CentralBankTask/internal/Interface"
	errPkg "CentralBankTask/internal/Middleware/Error"
	"github.com/labstack/echo"
)

// BankAPI structure of Bank API
type BankAPI struct {
	BankApplication Interface.BankInfoApplication
	Logger          errPkg.MultiLogger
}

// GetBankInfoHandler implementation of getting info
func (b BankAPI) GetBankInfoHandler(c echo.Context) error {
	//TODO implement me
	panic("implement me")
}

// NewUserHandler will initialize the articles/ resources endpoint
func NewUserHandler(e *echo.Echo, handler Interface.BankInfoAPI) {
	e.GET("/:date", handler.GetBankInfoHandler)
}
