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

func (h *handler) NewCapitalAccountHandlerFunc(c echo.Context) error {
	return view.HTMXNewCapital().Render(c.Request().Context(), c.Response())
}

func (h *handler) NewDebtAccountHandlerFunc(c echo.Context) error {
	return view.HTMXNewDebt().Render(c.Request().Context(), c.Response())
}

func (h *handler) NewExternalAccountHandlerFunc(c echo.Context) error {
	return view.HTMXNewExternal().Render(c.Request().Context(), c.Response())
}
