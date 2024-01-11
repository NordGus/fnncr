package main

import (
	"context"
	"log"
	"os"

	"github.com/NordGus/fnncr/authentication"
	"github.com/NordGus/fnncr/repository/sessionrepo"
	"github.com/NordGus/fnncr/repository/usersrepo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	ctx := context.Background()

	sessionRepo, err := sessionrepo.NewRedisRepository(
		func(opts *sessionrepo.RedisOpts) { opts.Ctx = ctx },
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer sessionRepo.Close()

	usersRepo, err := usersrepo.NewPostgreSQLRepository(
		func(opts *usersrepo.PostgreSQLOpts) { opts.Ctx = ctx },
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer usersRepo.Close()

	auth := authentication.New(
		func(opts *authentication.Opts) { opts.SessionRepository = sessionRepo },
		func(opts *authentication.Opts) { opts.UserRepository = usersRepo },
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
