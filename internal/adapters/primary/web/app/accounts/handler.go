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
	// TODO(#2): Saving Goals are a completely different data structure from accounts but it depends on accounts,
	// specifically savings accounts. As the balance changes in the different savings account so does the completion in
	// each savings goals. This goals contains an priority value order so the lower the priority value the earlier the
	// the goal is fullfil.
	SavingGoalsAccountsRoute = "/accounts/saving_goals"
	// TODO(#3): Debt accounts are basically all loans or any other debt you can have.
	DebtAccountsRoute = "/accounts/debts"
	// TODO(#4): External are basically transaction categories but in this case are easier administer.
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
