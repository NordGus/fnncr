package accounts

import (
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	// NOTE: Personal accounts are all accounts that can be a simple bank account and a savings account. Savings
	// accounts balance go towards fulfilling Savings Goals.
	PersonalAccountsRoute = "/accounts/personal"
	// NOTE: Debt accounts are loans and credit lines, basically any kind of personal debt that the user has or is owed.
	// Credit accounts are a little bit weird because they count as debts and assets. Where the available credit
	// count as an asset and the expended credit count as debt.
	DebtAccountsRoute = "/accounts/debt"
	// NOTE: External are basically transaction categories but in this case are easier administer.
	ExternalAccountsRoute = "/accounts/external"
)

type Handler interface {
	PersonalAccountsHandlerFunc(c echo.Context) error
	DebtAccountsHandlerFunc(c echo.Context) error
	ExternalAccountsHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
