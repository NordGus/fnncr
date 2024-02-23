package authentication

import (
	"net/http"
	"time"

	view "financo/internal/adapters/primary/web/app/views/authentication"
	"financo/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

func (h *handler) SignInHandlerFunc(c echo.Context) error {
	var (
		username = c.FormValue("username")
		password = c.FormValue("password")

		req = authentication.SignInUserReq{
			Username: username,
			Password: password,
		}

		form = view.FormLogin{
			ActionURL: SignInRoute,
			Username:  username,
			Password:  password,
			Failed:    true,
		}

		cookie = &http.Cookie{
			Name:    sessionCookieName,
			Path:    "/",
			Domain:  "localhost",
			Expires: time.Now().Add(sessionCookieMaxAge),

			MaxAge:   int(sessionCookieMaxAge.Seconds()),
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
	)

	resp, err := h.api.SignIn(c.Request().Context(), req)
	if err != nil {
		c.Logger().Error(err)

		return view.SignIn(form).Render(c.Request().Context(), c.Response())
	}

	cookie.Value = resp.SessionID

	c.SetCookie(cookie)
	c.Response().Header().Set("HX-Redirect", "/")

	return c.NoContent(http.StatusOK)
}
