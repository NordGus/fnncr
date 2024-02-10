package applets

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) BudgetAppletHandlerFunc(c echo.Context) error {
	return view.NotImplemented(
		layoutData(getUser(c), "fnncr | budget", budget),
		BudgetAppletRoute,
	).Render(c.Request().Context(), c.Response())
}
