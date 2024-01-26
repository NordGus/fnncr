package transactions

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/application"
	"github.com/labstack/echo/v4"
)

func (h *handler) Applet(c echo.Context) error {
	return view.NotImplemented(AppletRoute).Render(c.Request().Context(), c.Response())
}
