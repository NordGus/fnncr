package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) DebtAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		{AccType: models.LoanAccount, DisplayName: "Car Loan"},
		{AccType: models.LoanAccount, DisplayName: "Loan to friendly business"},
		{AccType: models.CreditAccount, DisplayName: "Credit Card"},
		{AccType: models.CreditAccount, DisplayName: "Credit Line"},
	}

	return view.DebtAccounts(acc).Render(c.Request().Context(), c.Response())
}
