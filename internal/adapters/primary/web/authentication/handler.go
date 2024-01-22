package authentication

import (
	"github.com/NordGus/fnncr/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

type Service interface {
	LoginHandler(c echo.Context) error
	SignInHandler(c echo.Context) error

	AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc
}

type service struct {
	api authentication.API
}

func New(api authentication.API) Service {
	return &service{
		api: api,
	}
}
