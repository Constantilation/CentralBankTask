package Interface

import "github.com/labstack/echo"

// BankInfoAPI implementation of user API interface
type BankInfoAPI interface {
	GetBankInfoHandler(c echo.Context) error
}
