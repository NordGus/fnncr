package applets

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) RootAppletHandlerFunc(c echo.Context) error {
	return view.Root(layoutData(getUser(c), "fnncr", root)).Render(c.Request().Context(), c.Response())
}
