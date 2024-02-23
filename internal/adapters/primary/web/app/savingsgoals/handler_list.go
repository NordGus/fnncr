package savingsgoals

import (
	"financo/internal/adapters/primary/web/app/models"
	view "financo/internal/adapters/primary/web/app/views/savingsgoals"
	"github.com/labstack/echo/v4"
)

func (h *handler) ListSavingsGoalsHandlerFunc(c echo.Context) error {
	goals := []models.SavingsGoal{
		models.NewSavingsGoal("Starter Emergency Fund", 100000, 100000, models.EUR_),
		models.NewSavingsGoal("3 Months Emergency Fund", 177342, 900000, models.EUR_),
		models.NewSavingsGoal("6 Months Emergency Fund", 0, 900000, models.EUR_),
	}

	return view.HTMXList(goals).Render(c.Request().Context(), c.Response())
}
