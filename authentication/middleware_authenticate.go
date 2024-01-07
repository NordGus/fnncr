package authentication

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *Service) AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Redirect(http.StatusTemporaryRedirect, "/login")

		return next(c)
	}
}
