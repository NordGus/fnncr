package accounts

import (
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	AppletRoute = "/accounts"

	NormalAccountsRoute = "/accounts/normal"
	LoanAccountsRoute   = "/accounts/loans"
	// NOTE: Credit accounts are a little bit weird because they count as debts and assets. Where the available credit
	// count as an asset and the expended credit count as debt.
	CreditAccountsRoute = "/accounts/credits"
	// NOTE: Savings accounts are basically normal accounts but their balance go to fulfill Savings Goals.
	SavingsAccountsRoute = "/accounts/savings"
	// NOTE: Saving Goals are a completely different data structure from accounts but it depends on accounts,
	// specifically savings accounts. As the balance changes in the different savings account so does the completion in
	// each savings goals. This goals contains an priority value order so the lower the priority value the earlier the
	// the goal is fullfil.
	SavingsGoalsAccountsRoute = "/accounts/savings_goals"
	// NOTE: External are basically transaction categories but in this case are easier administer.
	ExternalAccountsRoute = "/accounts/external"
)

type Handler interface {
	AppletHandlerFunc(c echo.Context) error
	NormalAccountsHandlerFunc(c echo.Context) error
	LoanAccountsHandlerFunc(c echo.Context) error
	CreditAccountsHandlerFunc(c echo.Context) error
	SavingsAccountsHandlerFunc(c echo.Context) error
	SavingsGoalsHandlerFunc(c echo.Context) error
	ExternalAccountsHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
