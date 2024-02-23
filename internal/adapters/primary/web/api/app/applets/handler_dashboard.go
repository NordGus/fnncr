package applets

import (
	"financo/internal/adapters/primary/web/api/app/accounts"
	"financo/internal/adapters/primary/web/api/app/savingsgoals"
	view "financo/internal/adapters/primary/web/api/app/views/applets"
	"github.com/labstack/echo/v4"
)

func (h *handler) DashboardAppletHandlerFunc(c echo.Context) error {
	sections := []view.DashboardSection{
		{
			Title:    "capital",
			Id:       "capital",
			FetchURL: accounts.CapitalAccountsRoute,
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

	actions := []view.ActionButton{
		{
			Name: "add savings goal",
			URL:  "/savings_goals/new",
		},
		{
			Name: "add account",
			URL:  "/accounts/new",
		},
	}

	return view.Dashboard(
		layoutData(getUser(c), "financo | dashboard", dashboard),
		sections,
		actions,
	).Render(c.Request().Context(), c.Response())
}
