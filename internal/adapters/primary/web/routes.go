package web

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/authentication"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	views "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/layouts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) setRoutes() {
	auth := authentication.New(a.AuthAPI)

	a.echo.Use(middleware.Logger())

	a.echo.StaticFS("/dist", a.assetsFS)

	a.echo.GET(authentication.LoginRoute, auth.LoginHandlerFunc)
	a.echo.POST(authentication.SignInRoute, auth.SignInHandlerFunc)
	a.echo.GET(authentication.SignOutRoute, auth.SignOutHandlerFunc, auth.AuthorizeMiddleware)

	a.echo.GET("/", func(c echo.Context) error {
		usr := c.Get(authentication.CurrentUserCtxKey).(models.User)

		return views.Application("fnncr", usr).Render(c.Request().Context(), c.Response())
	}, auth.AuthorizeMiddleware)
}
