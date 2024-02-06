package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) PersonalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		{
			AccType:        models.NormalAccount,
			DisplayName:    "My Personal Account",
			CurrentBalance: 426900,
		},
		{
			AccType:        models.NormalAccount,
			DisplayName:    "My Freelancer Account",
			CurrentBalance: -20000,
		},
		{
			AccType:        models.SavingsAccount,
			DisplayName:    "My Savings Account 1",
			CurrentBalance: 6900,
		},
		{
			AccType:        models.SavingsAccount,
			DisplayName:    "My Savings Account 2",
			CurrentBalance: 14400,
		},
	}

	return view.PersonalAccounts(acc).Render(c.Request().Context(), c.Response())
}
