package budget

import (
	"github.com/labstack/echo/v4"
)

const (
	// Routes

	AppletRoute = "/budget"
)

type Handler interface {
	AppletHandlerFunc(c echo.Context) error
}

type handler struct {
}

func New() Handler {
	return &handler{}
}
