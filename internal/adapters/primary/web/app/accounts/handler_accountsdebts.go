package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) DebtAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.LoanAccount, models.IOweDebt, "Car Loan", -426999, 1000000),
		models.NewAccount(models.LoanAccount, models.IAmOwedDebt, "Loan to friendly business", 4269, 42000),
		models.NewAccount(models.CreditAccount, models.IOweDebt, "Credit Card", -177300, 300000),
		models.NewAccount(models.CreditAccount, models.IOweDebt, "Credit Line", 0, 250000),
	}

	return view.DebtAccounts(acc).Render(c.Request().Context(), c.Response())
}
