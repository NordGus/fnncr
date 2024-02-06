package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) DebtAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		{
			AccType:        models.LoanAccount,
			DisplayName:    "Car Loan",
			CurrentBalance: 426900,
			Limit:          1000000,
		},
		{
			AccType:        models.LoanAccount,
			DisplayName:    "Loan to friendly business",
			CurrentBalance: 19200,
			Limit:          42000,
		},
		{
			AccType:        models.CreditAccount,
			DisplayName:    "Credit Card",
			CurrentBalance: 177300,
			Limit:          300000,
		},
		{
			AccType:        models.CreditAccount,
			DisplayName:    "Credit Line",
			CurrentBalance: 0,
			Limit:          250000,
		},
	}

	return view.DebtAccounts(acc).Render(c.Request().Context(), c.Response())
}
