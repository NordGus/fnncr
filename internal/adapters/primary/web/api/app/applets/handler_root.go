package applets

import (
	view "financo/internal/adapters/primary/web/api/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) RootAppletHandlerFunc(c echo.Context) error {
	return view.Root(
		layoutData(getUser(c), "financo", root),
	).Render(c.Request().Context(), c.Response())
}
