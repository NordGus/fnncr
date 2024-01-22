package web

import (
	"net/http"

	"github.com/NordGus/fnncr/internal/adapters/primary/web/authentication"
	"github.com/labstack/echo/v4"
)

func (a *App) SetRoutes() {
	auth := authentication.New(a.AuthAPI)

	a.echo.StaticFS("/dist", a.AssetsFS)

	a.echo.GET("/login", auth.LoginHandlerFunc)
	a.echo.POST("/sign_in", auth.SignInHandlerFunc)

	a.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<div>Hello There</div>")
	}, auth.AuthorizeMiddleware)
}
