package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) ExternalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Market",
			CurrentBalance: 10000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Subscriptions",
			CurrentBalance: 5000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Rent",
			CurrentBalance: 90000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Gift",
			CurrentBalance: 7000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Health Care",
			CurrentBalance: 15000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Free Time",
			CurrentBalance: 2500,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Interest",
			CurrentBalance: 20000,
		},
		{
			AccType:        models.ExternalAccount,
			DisplayName:    "Lottery",
			CurrentBalance: 2000,
		},
	}

	return view.ExternalAccounts(acc).Render(c.Request().Context(), c.Response())
}
