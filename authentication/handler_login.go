package authentication

import (
	"github.com/NordGus/fnncr/authentication/views"
	"github.com/labstack/echo/v4"
)

func (s *Service) LoginHandler(c echo.Context) error {
	return views.Login().Render(c.Request().Context(), c.Response())
}
