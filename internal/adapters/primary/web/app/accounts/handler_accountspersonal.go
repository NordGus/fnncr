package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) PersonalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		{AccType: models.NormalAccount, DisplayName: "My Personal Account"},
	}

	return view.PersonalAccounts(acc).Render(c.Request().Context(), c.Response())
}
