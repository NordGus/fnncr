package transactions

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/shared"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/application"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
)

func (h *handler) AppletHandlerFunc(c echo.Context) error {
	ald := c.Get(shared.ALDContextKey).(layouts.ApplicationLayoutData)

	ald.Title = "fnncr | transactions"

	for i := 0; i < len(ald.NavItems); i++ {
		if ald.NavItems[i].Name == "transactions" {
			ald.NavItems[i].IsActive = true

			break
		}
	}

	return view.NotImplemented(ald, AppletRoute).Render(c.Request().Context(), c.Response())
}
