package authentication

import (
	"net/http"

	"financo/internal/adapters_old/primary/web/api/app/models"
	"financo/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

func (h *handler) AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie(sessionCookieName)
		if err != nil {
			return c.Redirect(http.StatusTemporaryRedirect, LoginRoute)
		}

		err = cookie.Valid()
		if err != nil {
			c.Logger().Error(err)

			return c.Redirect(http.StatusTemporaryRedirect, LoginRoute)
		}

		resp, err := h.api.Authenticate(
			c.Request().Context(),
			authentication.AuthenticateUserReq{SessionID: cookie.Value},
		)
		if err != nil {
			c.Logger().Error(err)

			c.Response().Header().Set("HX-Redirect", LoginRoute)

			return c.Redirect(http.StatusTemporaryRedirect, LoginRoute)
		}

		c.Set(CurrentUserCtxKey, models.User{ID: resp.User.ID, Username: resp.User.Username})

		return next(c)
	}
}
