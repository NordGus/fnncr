package authentication

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Service) AuthenticateMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(s.sessionCookieName)
		if err != nil {
			c.Logger().Print(fmt.Errorf("authentication: unauthorized (reason: %v)", err))

			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		err = cookie.Valid()
		if err != nil {
			c.Logger().Print(fmt.Errorf("authentication: unauthorized (reason: %v)", err))

			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		return next(c)
	}
}
