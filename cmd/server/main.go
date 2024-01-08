package main

import (
	"os"

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

	app.StaticFS("/dist", os.DirFS("./dist"))

	app.GET("/", func(c echo.Context) error {
		return home().Render(c.Request().Context(), c.Response())
	}, auth.AuthenticateMiddleware)

	app.GET("/login", auth.LoginHandler)
	app.POST("/authenticate", auth.AuthenticateHandler)

	app.Logger.Fatal(app.Start(":4269"))
}
