package accounts

import (
	"financo/internal/adapters/primary/web/api/app/models"
	view "financo/internal/adapters/primary/web/api/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) DebtAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.LoanAccount, "Car Loan", (-1000000 + 426999), -1000000, models.EUR),
		models.NewAccount(models.LoanAccount, "Loan to friendly business", (42000 - 4269), 42000, models.EUR),
		models.NewAccount(models.CreditAccount, "Credit Card", -177300, 300000, models.EUR),
		models.NewAccount(models.CreditAccount, "Credit Line", 0, 250000, models.EUR),
	}

	return view.HTMXList(acc).Render(c.Request().Context(), c.Response())
}
