package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) NormalAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.NormalAccount, "My Personal Account", 426900, 0, models.EUR),
		models.NewAccount(models.NormalAccount, "My Freelancer Account", -20000, 0, models.EUR),
	}

	return view.HTMXList(acc).Render(c.Request().Context(), c.Response())
}
