package authentication

import (
	"github.com/NordGus/fnncr/authentication/views"
	"github.com/labstack/echo/v4"
)

func (s Service) LoginHandler(c echo.Context) error {
	return views.Login(views.LoginForm{
		Action: "/authenticate",
	}).Render(c.Request().Context(), c.Response())
}
