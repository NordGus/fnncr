package applets

import (
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/accounts"
	"github.com/NordGus/fnncr/internal/adapters/primary/web/app/savingsgoals"
	view "github.com/NordGus/fnncr/internal/adapters/primary/web/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) DashboardAppletHandlerFunc(c echo.Context) error {
	sections := []view.DashboardSection{
		{
			Title:    "capital",
			Id:       "capital",
			FetchURL: accounts.PersonalAccountsRoute,
		},
		{
			Title:    "debt",
			Id:       "debt",
			FetchURL: accounts.DebtAccountsRoute,
		},
		{
			Title:    "income and expenses",
			Id:       "external",
			FetchURL: accounts.ExternalAccountsRoute,
		},
		{
			Title:    "savings goals",
			Id:       "savings_goals",
			FetchURL: savingsgoals.SavingsGoalsRoute,
		},
	}

	return view.Dashboard(
		layoutData(getUser(c), "fnncr | dashboard", dashboard),
		sections,
	).Render(c.Request().Context(), c.Response())
}
