package authentication

import (
	"net/http"
	"time"

	"github.com/NordGus/fnncr/authentication/views"
	"github.com/labstack/echo/v4"
)

func (s Service) AuthenticateHandler(c echo.Context) error {
	var (
		username = c.FormValue("username")
		password = c.FormValue("password")

		cookie = &http.Cookie{
			Name:    s.sessionCookieName,
			Path:    "/",
			Domain:  s.sessionCookieDomain,
			Expires: time.Now().Add(s.sessionCookieDuration),

			MaxAge:   int(s.sessionCookieDuration.Seconds()),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
	)

	session, err := s.authenticate(username, password)
	if err != nil {
		c.Logger().Printf("authentication: unauthorized (reason: %s)", err)

		return views.Authenticate(views.LoginForm{
			Action:   "/authenticate",
			Username: username,
			Password: password,
			Failed:   true,
		}).Render(c.Request().Context(), c.Response())
	}

	cookie.Value = session
	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/")
}
