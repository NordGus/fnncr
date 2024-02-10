package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) PersonalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.NormalAccount, models.NoneDebt, "My Personal Account", 426900, 0),
		models.NewAccount(models.NormalAccount, models.NoneDebt, "My Freelancer Account", -20000, 0),
		models.NewAccount(models.SavingsAccount, models.NoneDebt, "My Savings Account 1", 6900, 0),
		models.NewAccount(models.SavingsAccount, models.NoneDebt, "My Savings Account 2", 14400, 0),
	}

	return view.PersonalAccounts(acc).Render(c.Request().Context(), c.Response())
}
