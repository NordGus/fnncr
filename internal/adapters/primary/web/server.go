package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/NordGus/fnncr/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

var (
	//go:embed dist
	assetsFS embed.FS
)

type App struct {
	Server   *http.Server
	echo     *echo.Echo
	assetsFS fs.FS
	AuthAPI  authentication.API
}

func NewApp(opts ...AppOption) *App {
	a := &App{
		echo:     echo.New(),
		assetsFS: echo.MustSubFS(assetsFS, "dist"),
	}

	for _, applyOption := range opts {
		applyOption(a)
	}

	return a
}

func (a *App) Run() error {
	a.setRoutes()

	return a.echo.StartServer(a.Server)
}
