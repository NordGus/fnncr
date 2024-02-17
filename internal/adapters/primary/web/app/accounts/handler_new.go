package accounts

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) NewHandlerFunc(c echo.Context) error {
	options := []view.CreationOption{
		{
			Name:        "capital",
			Description: "any normal bank account or savings account",
			URL:         NewCapitalAccountRoute,
		},
		{
			Name:        "debt",
			Description: "any loan or credit line you have",
			URL:         NewDebtAccountRoute,
		},
		{
			Name:        "external",
			Description: "all sources of income or expenses",
			URL:         NewExternalAccountRoute,
		},
	}

	return view.HTMXNew(options).Render(c.Request().Context(), c.Response())
}
