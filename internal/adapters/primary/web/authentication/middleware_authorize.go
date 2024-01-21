package authentication

import "github.com/labstack/echo/v4"

func (h *Handler) AuthorizeMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		panic("unimplemented")
	}
}
