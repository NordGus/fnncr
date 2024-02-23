package applets

import (
	view "financo/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) IntelligenceAppletHandlerFunc(c echo.Context) error {
	return view.NotImplemented(
		layoutData(getUser(c), "fnncr | intelligence", intelligence),
		IntelligenceAppletRoute,
	).Render(c.Request().Context(), c.Response())
}
