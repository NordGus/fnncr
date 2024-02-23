package authentication

import (
	view "financo/internal/adapters/primary/web/app/views/authentication"
	"github.com/labstack/echo/v4"
)

func (h *handler) LoginHandlerFunc(c echo.Context) error {
	form := view.FormLogin{
		ActionURL: SignInRoute,
	}

	return view.Login(form).Render(c.Request().Context(), c.Response())
}
