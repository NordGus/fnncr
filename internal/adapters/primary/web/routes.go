package web

import (
	"net/http"

	"github.com/NordGus/fnncr/internal/adapters/primary/web/authentication"
	"github.com/labstack/echo/v4"
)

func (a *App) SetRoutes() {
	authHandler := authentication.New(a.AuthAPI)

	a.echo.GET("/login", authHandler.LogInHandlerFunc)
	a.echo.POST("/sign_in", authHandler.SignInHandlerFunc)

	a.echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<div>Hello There</div>")
	}, authHandler.AuthorizeMiddleware)
}
