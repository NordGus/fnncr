package main

import (
	"github.com/NordGus/fnncr/authentication"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	var (
		app = echo.New()

		auth = authentication.New()
	)

	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		return home().Render(c.Request().Context(), c.Response())
	}, auth.AuthenticateMiddleware)

	app.GET("/login", auth.LoginHandler)

	app.Logger.Fatal(app.Start(":4269"))
}
