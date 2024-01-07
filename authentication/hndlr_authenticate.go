package authentication

import (
	"net/http"
	"time"

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

	c.Logger().Printf("%s %s", username, password)

	// TODO: Implement a data storage to retrieve users and authenticate them

	// TODO: Implement a mechanism to render the form and return it to indicate
	//   that it failed to access the application

	c.SetCookie(cookie)

	return c.Redirect(http.StatusFound, "/")
}
