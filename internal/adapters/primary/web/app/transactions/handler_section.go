package transactions

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/shared"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/application"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
)

func (h *handler) AppletHandlerFunc(c echo.Context) error {
	ald := c.Get(shared.ALDContextKey).(layouts.ApplicationLayoutData)

	return view.NotImplemented(ald, AppletRoute).Render(c.Request().Context(), c.Response())
}
