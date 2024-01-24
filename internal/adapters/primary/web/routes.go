package web

import (
	"net/http"

	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/authentication"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (a *App) setRoutes() {
	auth := authentication.New(a.AuthAPI)

	a.echo.Use(middleware.Logger())

	a.echo.StaticFS("/dist", a.assetsFS)

	a.echo.GET(authentication.LoginRoute, auth.LoginHandlerFunc)
	a.echo.POST(authentication.SignInRoute, auth.SignInHandlerFunc)

	a.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<div>Hello There</div>")
	}, auth.AuthorizeMiddleware)
}
