package accounts

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) SavingGoalsHandlerFunc(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}
