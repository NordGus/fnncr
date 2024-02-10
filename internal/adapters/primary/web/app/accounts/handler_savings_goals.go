package accounts

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/models"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/accounts"
	"github.com/labstack/echo/v4"
)

func (h *handler) SavingsGoalsHandlerFunc(c echo.Context) error {
	goals := []models.SavingsGoal{
		{GoalName: "Goal 1", Goal: 177300, Saved: 177300},
		{GoalName: "Goal 2", Goal: 1000000, Saved: 177342},
	}

	return view.SavingGoals(goals).Render(c.Request().Context(), c.Response())
}
