package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) ExternalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Market", 10000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Subscriptions", 5000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Rent", 90000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Gifts", 7000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Health Care", 15000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Free Time", 2500, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Interest", 20000, 0),
		models.NewAccount(models.ExternalAccount, models.NoneDebt, "Lottery", 2000, 0),
	}

	return view.ExternalAccounts(acc).Render(c.Request().Context(), c.Response())
}
