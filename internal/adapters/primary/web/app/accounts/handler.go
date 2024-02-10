package accounts

import (
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	NormalAccountsRoute = "/accounts/normal"
	LoanAccountsRoute   = "/accounts/loans"
	// NOTE: Credit accounts are a little bit weird because they count as debts and assets. Where the available credit
	// count as an asset and the expended credit count as debt.
	CreditAccountsRoute = "/accounts/credits"
	// NOTE: Savings accounts are basically normal accounts but their balance go to fulfill Savings Goals.
	SavingsAccountsRoute = "/accounts/savings"
	// NOTE: External are basically transaction categories but in this case are easier administer.
	ExternalAccountsRoute = "/accounts/external"
)

type Handler interface {
	NormalAccountsHandlerFunc(c echo.Context) error
	LoanAccountsHandlerFunc(c echo.Context) error
	CreditAccountsHandlerFunc(c echo.Context) error
	SavingsAccountsHandlerFunc(c echo.Context) error
	ExternalAccountsHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
