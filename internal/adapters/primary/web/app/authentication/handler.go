package authentication

import (
	"time"

	"financo/internal/core/services/authentication"
	"github.com/labstack/echo/v4"
)

const (
	CurrentUserCtxKey = "currentUser"

	sessionCookieName   = "_session_financo"
	sessionCookieMaxAge = 30 * 24 * time.Hour

	// Routes

	LoginRoute   = "/login"
	SignInRoute  = "/sign_in"
	SignOutRoute = "/sign_out"
)

type Handler interface {
	LoginHandlerFunc(c echo.Context) error
	SignInHandlerFunc(c echo.Context) error
	SignOutHandlerFunc(c echo.Context) error

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
