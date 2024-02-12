package savingsgoals

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/savingsgoals"
	"github.com/labstack/echo/v4"
)

func (h *handler) ListSavingsGoalsHandlerFunc(c echo.Context) error {
	goals := []models.SavingsGoal{
		models.NewSavingsGoal("Starter Emergency Fund", 100000, 100000),
		models.NewSavingsGoal("3 Months Emergency Fund", 177342, 900000),
		models.NewSavingsGoal("6 Months Emergency Fund", 0, 900000),
	}

	return view.HTMXList(goals).Render(c.Request().Context(), c.Response())
}
