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

		record, err := s.getCurrentUser(cookie)
		if err != nil {
			c.Logger().Print(fmt.Errorf("authentication: unauthorized (reason: %v)", err))

			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		c.Set("user", record)

		return next(c)
	}
}

func (s Service) getCurrentUser(cookie *http.Cookie) (UserRecord, error) {
	err := cookie.Valid()
	if err != nil {
		return nil, err
	}

	session, err := s.sessionRepository.Get(cookie.Value)
	if err != nil {
		return nil, err
	}

	record, err := s.userRepository.GetByID(session.UserID())
	if err != nil {
		return nil, err
	}

	return record, nil
}
