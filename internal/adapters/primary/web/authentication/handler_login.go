package authentication

import (
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/views/authentication"
	"github.com/labstack/echo/v4"
)

func (h *service) LoginHandler(c echo.Context) error {
	form := view.FormLogin{
		ActionURL: "/sign_in",
	}

	return view.Login(form).Render(c.Request().Context(), c.Response())
}
