package authentication

import (
	"github.com/NordGus/fnncr/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

type Handler interface {
	LoginHandlerFunc(c echo.Context) error
	SignInHandlerFunc(c echo.Context) error

	AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type handler struct {
	api authentication.API
}

func New(api authentication.API) Handler {
	return &handler{
		api: api,
	}
}
