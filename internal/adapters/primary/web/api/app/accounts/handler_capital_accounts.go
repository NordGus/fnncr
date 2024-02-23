package accounts

import (
	"financo/internal/adapters/primary/web/api/app/models"
	view "financo/internal/adapters/primary/web/api/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) CapitalHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.NormalAccount, "My Personal Account", 426900, 0, models.EUR),
		models.NewAccount(models.NormalAccount, "My Freelancer Account", -20000, 0, models.EUR),
		models.NewAccount(models.SavingsAccount, "My Savings Account 1", 6900, 0, models.EUR),
		models.NewAccount(models.SavingsAccount, "My Savings Account 2", 14400, 0, models.EUR),
	}

	return view.HTMXList(acc).Render(c.Request().Context(), c.Response())
}
