package web

import (
	"net/http"

	"github.com/NordGus/fnncr/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

type App struct {
	Server  *http.Server
	echo    *echo.Echo
	AuthAPI authentication.API
}

func NewApp(opts ...AppOption) *App {
	a := &App{
		echo: echo.New(),
	}

	for _, applyOption := range opts {
		applyOption(a)
	}

	return a
}

func (a *App) Run() error {
	return a.echo.StartServer(a.Server)
}
