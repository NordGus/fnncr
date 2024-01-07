package main

import (
	"context"

	"github.com/NordGus/fnncr/authentication"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var (
		ctx, cancel = context.WithCancelCause(context.Background())
		app         = echo.New()

		auth = authentication.New(ctx, cancel)
	)

	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return home().Render(c.Request().Context(), c.Response())
	}, auth.AuthenticateMiddleware)

	app.GET("/login", auth.LoginHandler)

	app.Logger.Fatal(app.Start(":4269"))
}
