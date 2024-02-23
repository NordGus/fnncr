package applets

import (
	view "financo/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) BookAppletHandlerFunc(c echo.Context) error {
	return view.NotImplemented(
		layoutData(getUser(c), "fnncr | book", book),
		BookAppletRoute,
	).Render(c.Request().Context(), c.Response())
}
