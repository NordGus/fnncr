package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) LoanAccountsHandlerFunc(c echo.Context) error {
	acc := []models.Account{
		models.NewAccount(models.LoanAccount, "Car Loan", (-1000000 + 426999), -1000000, models.EUR),
		models.NewAccount(models.LoanAccount, "Loan to friendly business", (42000 - 4269), 42000, models.EUR),
	}

	return view.HTMXList(acc).Render(c.Request().Context(), c.Response())
}
