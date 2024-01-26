package accounts

import (
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	AppletRoute = "/accounts"

	// TODO(#0): All accounts can have a parent, this is used to create categories each one.
	// TODO(#1): Personal accounts convey bank accounts, savings accounts and credit lines/cards. Credit lines/cards
	// can be added as expenses for easier cataloging of payments in the application and following of personal expenses.
	PersonalAccountsRoute = "/accounts/personal"
	// TODO(#2): Debt accounts are basically all loans or any other debt you can have.
	DebtAccountsRoute = "/accounts/debts"
	// TODO(#3): External are basically transaction categories but in this case are easier administer.
	ExternalAccountsRoute = "/accounts/external"
)

type Handler interface {
	AppletHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
