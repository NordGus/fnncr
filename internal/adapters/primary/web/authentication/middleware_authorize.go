package authentication

import (
	"net/http"

	"github.com/NordGus/fnncr/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

func (h *Handler) AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("_session_fnncr")
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		err = cookie.Valid()
		if err != nil {
			c.Logger().Error(err)

			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		resp, err := h.api.AuthenticateSession(
			c.Request().Context(),
			authentication.AuthenticateUserReq{SessionID: cookie.Value},
		)
		if err != nil {
			c.Logger().Error(err)

			return c.Redirect(http.StatusTemporaryRedirect, "/login")
		}

		c.Logger().Debug(resp.User)

		c.Set("currentUser", resp.User)

		return next(c)
	}
}
