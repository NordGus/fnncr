package applets

import (
	view "financo/internal/adapters_old/primary/web/api/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) BudgetAppletHandlerFunc(c echo.Context) error {
	return view.NotImplemented(
		layoutData(getUser(c), "financo | budget", budget),
		BudgetAppletRoute,
	).Render(c.Request().Context(), c.Response())
}
